package authorization

import (
	"fmt"
	"net/http"
	"pubapi/internal/model"
)

func (ctl *Controller) Deal() {
	// 参数校验
	if ctl.param.Code == "" {
		ctl.NewErrorResponse(http.StatusBadRequest, "missing request parameters")
		return
	}
	authCode := model.NewAuthorizationCode(ctl.param.Code)
	session, err := authCode.MakeSession()
	if err != nil {
		ctl.NewErrorResponse(http.StatusInternalServerError, "wechat server error")
	}
	fmt.Printf("%s: %+v\n", ctl.param.Code, session)
	ctl.NewOkResponse(http.StatusOK, Reply{Message: "success"})
}
