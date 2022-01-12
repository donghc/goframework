package interceptor

import "goframework/internal/pkg/core"

func (i *interceptor) CheckRBAC() core.HandlerFunc {
	return func(c core.Context) {

	}
}
