package core

import "goframework/internal/proposal"

type Option func(*option)

type option struct {
	disablePProf      bool
	disableSwagger    bool
	disablePrometheus bool
	enableCors        bool
	enableRate        bool
	enableOpenBrowser string
	alertNotify       proposal.NotifyHandler
	recordHandler     proposal.RecordHandler
}


// WithDisablePProf 禁用 pprof
func WithDisablePProf() Option {
	return func(opt *option) {
		opt.disablePProf = true
	}
}

// WithDisableSwagger 禁用 swagger
func WithDisableSwagger() Option {
	return func(opt *option) {
		opt.disableSwagger = true
	}
}

// WithDisablePrometheus 禁用prometheus
func WithDisablePrometheus() Option {
	return func(opt *option) {
		opt.disablePrometheus = true
	}
}

// WithAlertNotify 设置告警通知
func WithAlertNotify(notifyHandler proposal.NotifyHandler) Option {
	return func(opt *option) {
		opt.alertNotify = notifyHandler
	}
}

// WithRecordMetrics 设置记录接口指标
func WithRecordMetrics(recordHandler proposal.RecordHandler) Option {
	return func(opt *option) {
		opt.recordHandler = recordHandler
	}
}

// WithEnableOpenBrowser 启动后在浏览器中打开 uri
func WithEnableOpenBrowser(uri string) Option {
	return func(opt *option) {
		opt.enableOpenBrowser = uri
	}
}

// WithEnableCors 设置支持跨域
func WithEnableCors() Option {
	return func(opt *option) {
		opt.enableCors = true
	}
}

// WithEnableRate 设置支持限流
func WithEnableRate() Option {
	return func(opt *option) {
		opt.enableRate = true
	}
}
