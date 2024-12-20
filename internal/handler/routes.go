// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	file "github.com/colinrs/protohub/internal/handler/file"
	project "github.com/colinrs/protohub/internal/handler/project"
	role "github.com/colinrs/protohub/internal/handler/role"
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
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/project",
				Handler: project.GetProjectByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/project/create",
				Handler: project.CreateProjectHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/project/delete",
				Handler: project.DeleteProjectHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/project/list",
				Handler: project.GetProjectListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/project/project_user_list",
				Handler: project.GetProjectUserListHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/project/update",
				Handler: project.UpdateProjectHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/role",
				Handler: role.GetRoleByCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/role/list",
				Handler: role.GetRoleListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user",
				Handler: user.GetUserByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/change_password",
				Handler: user.ChangePasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/create",
				Handler: user.CreateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/user/delete",
				Handler: user.DeleteUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/join_project",
				Handler: user.JoinProjectHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/leave_project",
				Handler: user.LeaveProjectHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/list",
				Handler: user.GetUserListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/logout",
				Handler: user.LogoutHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/project/list",
				Handler: user.GetUserProjectListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/refresh_token",
				Handler: user.RefreshTokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/update",
				Handler: user.UpdateUserHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
