package router

import (
	v1 "pubapi/api/v1"

	"github.com/gin-gonic/gin"
)

const wxapp_path = "wxapp"

func authorization(router *gin.RouterGroup) {
	router.GET("/authorization", v1.WxappAuthorization)
}
