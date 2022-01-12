package interceptor

import (
	"goframework/internal/pkg/core"
	"goframework/internal/proposal"
)

func (i *interceptor) CheckLogin(ctx core.Context) (sessionUserInfo proposal.SessionUserInfo, err core.BusinessError) {

	return proposal.SessionUserInfo{}, nil
}
