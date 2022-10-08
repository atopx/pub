package authorization

import (
	"pubapi/internal/controller"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	param *Param
	controller.Controller
}

func NewController(ctx *gin.Context) (*Controller, error) {
	ctl := Controller{param: new(Param)}
	ctl.ContextLoader(ctx)
	if err := ctl.NewRequestParam(ctl.param); err != nil {
		return nil, err
	}
	return &ctl, nil
}

type Param struct {
	Code  string `form:"code"`
	Name  string `form:"name"`
	Image string `form:"image"`
}

type Reply struct {
	Message string `json:"message"`
}
