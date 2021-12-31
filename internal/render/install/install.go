package install

import (
	"goframework/configs"
	"goframework/internal/pkg/core"
	"goframework/pkg/file"
	"net/http"
	"runtime"

	"go.uber.org/zap"
)

type handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) View() core.HandlerFunc {
	type viewResponse struct {
		Config       configs.Config
		MinGoVersion float64
		GoVersion    string
	}

	return func(ctx core.Context) {
		if _, ok := file.IsExists(configs.ProjectInstallMark); ok {
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
		}

		obj := new(viewResponse)
		obj.Config = configs.Get()
		obj.MinGoVersion = configs.MinGoVersion
		obj.GoVersion = runtime.Version()
		ctx.HTML("install_view", obj)
	}
}