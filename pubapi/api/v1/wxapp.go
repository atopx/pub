package v1

import (
	"pubapi/internal/controller/wxapp/authorization"

	"github.com/gin-gonic/gin"
)

// WxappAuthorization
// @summary 小程序认证接口
// @Accept json
// @Produce json
// @Param param query authorization.Param true "请求参数"
// @Response 200 object authorization.Reply "调用成功"
// @Response 400 object system.ReplyError "请求错误"
// @Response 500 object system.ReplyError "系统错误"
// @Router /api/v1/wxapp/authorization [get]
func WxappAuthorization(ctx *gin.Context) {
	if controller, err := authorization.NewController(ctx); err == nil {
		controller.Deal()
	}
}
