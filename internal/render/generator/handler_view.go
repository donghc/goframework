package generator_handler

import (
	"goframework/internal/pkg/core"
)

func (h *handler) HandlerView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("generator_handler", nil)
	}
}
