// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"love_mall/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/sms/code",
				Handler: SendCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login/wx",
				Handler: WxLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login/phone",
				Handler: PhoneLoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/edit",
				Handler: EditInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/user"),
	)
}
