// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	file "github.com/colinrs/protohub/internal/handler/file"
	user "github.com/colinrs/protohub/internal/handler/user"
	"github.com/colinrs/protohub/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: file.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/download",
				Handler: file.DownloadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: file.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: file.UploadHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/file"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/detail",
					Handler: user.DetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/refresh/authorization",
					Handler: user.RefreshAuthorizationHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/mail/code/send/register",
				Handler: user.MailCodeSendRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/user"),
	)
}
