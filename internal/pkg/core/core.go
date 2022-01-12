package core

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	cors "github.com/rs/cors/wrapper/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"goframework/assets"
	"goframework/configs"
	"goframework/internal/code"
	"goframework/internal/proposal"
	"goframework/pkg/browser"
	"goframework/pkg/env"
	"goframework/pkg/errors"
	"goframework/pkg/trace"
	"golang.org/x/time/rate"
	"html/template"
	"net/http"
	"net/url"
	"runtime/debug"
	"time"
)

const _UI = `
                          $$\                           
                          $$ |                          
$$$$$$$\   $$$$$$\        $$$$$$$\  $$\   $$\  $$$$$$\  
$$  __$$\ $$  __$$\       $$  __$$\ $$ |  $$ |$$  __$$\ 
$$ |  $$ |$$ /  $$ |      $$ |  $$ |$$ |  $$ |$$ /  $$ |
$$ |  $$ |$$ |  $$ |      $$ |  $$ |$$ |  $$ |$$ |  $$ |
$$ |  $$ |\$$$$$$  |      $$$$$$$  |\$$$$$$  |\$$$$$$$ |
\__|  \__| \______/       \_______/  \______/  \____$$ |
                                              $$\   $$ |
                                              \$$$$$$  |
                                               \______/
`

// DisableTraceLog 禁止记录日志
func DisableTraceLog(ctx Context) {
	ctx.disableTrace()
}

// DisableRecordMetrics 禁止记录指标
func DisableRecordMetrics(ctx Context) {
	ctx.disableRecordMetrics()
}

// AliasForRecordMetrics 对请求路径起个别名，用于记录指标。
// 如：Get /user/:username 这样的路径，因为 username 会有非常多的情况，这样记录指标非常不友好。
func AliasForRecordMetrics(path string) HandlerFunc {
	return func(ctx Context) {
		ctx.setAlias(path)
	}
}

// WrapAuthHandler 用来处理 Auth 的入口
func WrapAuthHandler(handler func(Context) (sessionUserInfo proposal.SessionUserInfo, err BusinessError)) HandlerFunc {
	return func(ctx Context) {
		sessionUserInfo, err := handler(ctx)
		if err != nil {
			ctx.AbortWithError(err)
			return
		}

		ctx.setSessionUserInfo(sessionUserInfo)
	}
}

var _ Mux = (*mux)(nil)

// Mux http mux
type Mux interface {
	http.Handler
	Group(relativePath string, handlers ...HandlerFunc) RouterGroup
}

type mux struct {
	engine *gin.Engine
}

func (m *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m.engine.ServeHTTP(w, req)
}

func (m *mux) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	return &router{
		group: m.engine.Group(relativePath, wrapHandlers(handlers...)...),
	}
}

func New(logger *zap.Logger, options ...Option) (Mux, error) {
	if logger == nil {
		return nil, errors.New("logger must be required")
	}
	gin.SetMode(gin.ReleaseMode)
	mux := &mux{
		engine: gin.New(),
	}
	//fmt.Println(color.Blue(_UI))
	//设置静态资源文件
	mux.engine.StaticFS("assets", http.FS(assets.Bootstrap))
	//加载html页面
	mux.engine.SetHTMLTemplate(template.Must(template.New("").ParseFS(assets.Templates, "templates/**/*")))

	// withoutTracePaths 这些请求，默认不记录日志
	withoutTracePaths := map[string]bool{
		"/metrics":                  true,
		"/debug/pprof/":             true,
		"/debug/pprof/cmdline":      true,
		"/debug/pprof/profile":      true,
		"/debug/pprof/symbol":       true,
		"/debug/pprof/trace":        true,
		"/debug/pprof/allocs":       true,
		"/debug/pprof/block":        true,
		"/debug/pprof/goroutine":    true,
		"/debug/pprof/heap":         true,
		"/debug/pprof/mutex":        true,
		"/debug/pprof/threadcreate": true,
		"/favicon.ico":              true,
		"/system/health":            true,
	}
	//加载options选项
	opt := new(option)
	for _, f := range options {
		f(opt)
	}
	if !opt.disablePProf {
		//生产环境不开放pprof
		if !env.Active().IsPro() {
			pprof.Register(mux.engine) //register pprof to gin
		}
	}

	if !opt.disableSwagger {
		//生产环境不开放swagger
		if !env.Active().IsPro() {
			mux.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //注册swagger
		}
	}

	if !opt.disablePrometheus {
		mux.engine.GET("/metrics", gin.WrapH(promhttp.Handler())) //register prometheus
	}

	//允许跨域请求
	if opt.enableCors {
		mux.engine.Use(cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedHeaders:     []string{"*"},
			AllowCredentials:   true,
			OptionsPassthrough: true,
		}))
	}

	if opt.enableOpenBrowser != "" {
		err := browser.Open(opt.enableOpenBrowser)
		fmt.Println("", err)
	}

	// recover两次，防止处理时发生panic，尤其是在OnPanicNotify中。
	mux.engine.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", string(debug.Stack())))
			}
		}()
		ctx.Next()
	})

	mux.engine.Use(func(ctx *gin.Context) {
		if ctx.Writer.Status() == http.StatusNotFound {
			return
		}
		ts := time.Now()
		context := newContext(ctx)
		defer releaseContext(context)
		context.init()
		context.setLogger(logger)
		context.ableRecordMetrics()

		if !withoutTracePaths[ctx.Request.URL.Path] {
			if traceId := context.GetHeader(trace.Header); traceId != "" {
				context.setTrace(trace.New(traceId))
			} else {
				context.setTrace(trace.New(""))
			}
		}
		defer func() {
			var (
				response        interface{}
				businessCode    int
				businessCodeMsg string
				abortErr        error
				traceId         string
			)

			if ct := context.Trace(); ct != nil {
				context.SetHeader(trace.Header, ct.ID())
				traceId = ct.ID()
			}
			// region 发生 Panic 异常发送告警提醒
			if err := recover(); err != nil {
				stackInfo := string(debug.Stack())
				logger.Error("got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", stackInfo))
				context.AbortWithError(Error(
					http.StatusInternalServerError,
					code.ServerError,
					code.Text(code.ServerError)),
				)
				if notifyHandler := opt.alertNotify; notifyHandler != nil {
					notifyHandler(&proposal.AlertMessage{
						ProjectName:  configs.ProjectName,
						Env:          env.Active().Value(),
						TraceID:      traceId,
						HOST:         context.Host(),
						URI:          context.URI(),
						Method:       context.Method(),
						ErrorMessage: err,
						ErrorStack:   stackInfo,
						Timestamp:    time.Now()},
					)
				}
			}
			//endregion

			// region 发生错误，进行返回
			if ctx.IsAborted() {
				for i := range ctx.Errors {
					multierr.AppendInto(&abortErr, ctx.Errors[i])
				}
				if err := context.abortError(); err != nil {
					//判断是否需要发送告警通知
					if err.IsAlert() {
						if notifyHandler := opt.alertNotify; notifyHandler != nil {
							notifyHandler(&proposal.AlertMessage{
								ProjectName:  configs.ProjectName,
								Env:          env.Active().Value(),
								TraceID:      traceId,
								HOST:         context.Host(),
								URI:          context.URI(),
								Method:       context.Method(),
								ErrorMessage: err.Message(),
								ErrorStack:   fmt.Sprintf("%+v", err.StackError()),
								Timestamp:    time.Now(),
							})
						}
					}

					multierr.AppendInto(&abortErr, err.StackError())
					businessCode = err.BusinessCode()
					businessCodeMsg = err.Message()
					response = &code.Failure{
						Code:    businessCode,
						Message: businessCodeMsg,
					}
					ctx.JSON(err.HTTPCode(), response)
				}
			}
			//endregion

			// region 正确返回
			response = context.getPayload()
			if response != nil {
				ctx.JSON(http.StatusOK, response)
			}
			//endregion

			// region 记录指标
			if opt.recordHandler != nil && context.isRecordMetrics() {
				path := context.Path()
				if alias := context.Alias(); alias != "" {
					path = alias
				}

				opt.recordHandler(&proposal.MetricsMessage{
					ProjectName:  configs.ProjectName,
					Env:          env.Active().Value(),
					TraceID:      traceId,
					HOST:         context.Host(),
					Path:         path,
					Method:       context.Method(),
					HTTPCode:     ctx.Writer.Status(),
					BusinessCode: businessCode,
					CostSeconds:  time.Since(ts).Seconds(),
					IsSuccess:    !ctx.IsAborted() && (ctx.Writer.Status() == http.StatusOK),
				})
			}
			//endregion

			// region 记录日志
			var t *trace.Trace
			if x := context.Trace(); x != nil {
				t = x.(*trace.Trace)
			} else {
				return
			}

			decodedURL, _ := url.QueryUnescape(ctx.Request.URL.RequestURI())

			// ctx.Request.Header，精简 Header 参数
			traceHeader := map[string]string{
				"Content-Type":              ctx.GetHeader("Content-Type"),
				configs.HeaderLoginToken:    ctx.GetHeader(configs.HeaderLoginToken),
				configs.HeaderSignToken:     ctx.GetHeader(configs.HeaderSignToken),
				configs.HeaderSignTokenDate: ctx.GetHeader(configs.HeaderSignTokenDate),
			}

			t.WithRequest(&trace.Request{
				TTL:        "un-limit",
				Method:     ctx.Request.Method,
				DecodedURL: decodedURL,
				Header:     traceHeader,
				Body:       string(context.RawData()),
			})

			var responseBody interface{}

			if response != nil {
				responseBody = response
			}

			t.WithResponse(&trace.Response{
				Header:          ctx.Writer.Header(),
				HttpCode:        ctx.Writer.Status(),
				HttpCodeMsg:     http.StatusText(ctx.Writer.Status()),
				BusinessCode:    businessCode,
				BusinessCodeMsg: businessCodeMsg,
				Body:            responseBody,
				CostSeconds:     time.Since(ts).Seconds(),
			})

			t.Success = !ctx.IsAborted() && (ctx.Writer.Status() == http.StatusOK)
			t.CostSeconds = time.Since(ts).Seconds()

			logger.Info("trace-log",
				zap.Any("method", ctx.Request.Method),
				zap.Any("path", decodedURL),
				zap.Any("http_code", ctx.Writer.Status()),
				zap.Any("business_code", businessCode),
				zap.Any("success", t.Success),
				zap.Any("cost_seconds", t.CostSeconds),
				zap.Any("trace_id", t.Identifier),
				zap.Any("trace_info", t),
				zap.Error(abortErr),
			)
			// endregion

		}()
		ctx.Next()

	})

	//开启限流
	if opt.enableRate {
		limiter := rate.NewLimiter(rate.Every(time.Second*1), configs.MaxRequestsPerSecond)
		mux.engine.Use(func(ctx *gin.Context) {
			context := newContext(ctx)
			defer releaseContext(context)
			if !limiter.Allow() {
				context.AbortWithError(Error(
					http.StatusTooManyRequests,
					code.TooManyRequests,
					code.Text(code.TooManyRequests)),
				)
				return
			}
		})
	}

	mux.engine.NoMethod(wrapHandlers(DisableTraceLog)...)
	mux.engine.NoRoute(wrapHandlers(DisableTraceLog)...)

	system := mux.Group("/system")
	{
		//健康检查
		system.GET("/health", func(ctx Context) {
			resp := &struct {
				Timestamp   time.Time `json:"timestamp"`
				Environment string    `json:"environment"`
				Host        string    `json:"host"`
				Status      string    `json:"status"`
			}{
				Timestamp:   time.Now(),
				Environment: env.Active().Value(),
				Host:        ctx.Host(),
				Status:      "ok",
			}
			ctx.Payload(resp)
		})
	}

	return mux, nil
}
