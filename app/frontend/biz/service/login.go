package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/juzi0911/gomall_PJ/app/frontend/hertz_gen/frontend/auth"
	common "github.com/juzi0911/gomall_PJ/app/frontend/hertz_gen/frontend/common"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	session.Save()

	return
}
