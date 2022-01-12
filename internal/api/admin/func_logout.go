package admin

import "goframework/internal/pkg/core"

type logoutResponse struct {
	Username string `json:"username"` // 用户账号
}

// Logout 管理员登出
// @Summary 管理员登出
// @Description 管理员登出
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} logoutResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/logout [post]
// @Security LoginToken
func (h *handler) Logout() core.HandlerFunc {
	return func(c core.Context) {

	}
}
