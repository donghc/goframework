package config

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"goframework/configs"
	"goframework/internal/code"
	"goframework/internal/pkg/core"
	"goframework/pkg/env"
	"goframework/pkg/mail"
	"net/http"
)

type emailRequest struct {
	Host string `form:"host"` // 邮箱服务器
	Port string `form:"port"` // 端口
	User string `form:"user"` // 发件人邮箱
	Pass string `form:"pass"` // 发件人密码
	To   string `form:"to"`   // 收件人邮箱地址，多个用,分割
}

type emailResponse struct {
	Email string `json:"email"` // 邮箱地址
}

func (h *handler) Email() core.HandlerFunc {
	return func(c core.Context) {
		req := new(emailRequest)
		resp := new(emailResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		options := &mail.Options{
			MailHost: req.Host,
			MailPort: cast.ToInt(req.Port),
			MailUser: req.User,
			MailPass: req.Pass,
			MailTo:   req.To,
			Subject:  fmt.Sprintf("%s[%s] 邮箱告警人调整通知。", configs.ProjectName, env.Active().Value()),
			Body:     fmt.Sprintf("%s[%s] 已添加您为系统告警通知人。", configs.ProjectName, env.Active().Value()),
		}

		if err := mail.Send(options); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SendEmailError,
				code.Text(code.SendEmailError)).WithError(err),
			)
		}
		//更新配置文件
		viper.Set("mail.host", req.Host)
		viper.Set("mail.port", cast.ToInt(req.Port))
		viper.Set("mail.user", req.User)
		viper.Set("mail.pass", req.Pass)
		viper.Set("mail.to", req.To)

		err := viper.WriteConfig()
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WriteConfigError,
				code.Text(code.WriteConfigError)).WithError(err),
			)
			return
		}
		resp.Email = req.To
		c.Payload(resp)
	}
}
