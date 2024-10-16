package project

import (
	"context"

	"github.com/colinrs/protohub/internal/manage"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjecUsertListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	projectManage manage.ProjectManage
}

func NewGetProjecUsertListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjecUsertListLogic {
	return &GetProjecUsertListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		projectManage: manage.NewProjectManage(ctx, svcCtx),
	}
}

func (l *GetProjecUsertListLogic) GetProjecUsertList(req *types.GetProjectUserListRequest) (resp *types.GetProjectUserListResponse, err error) {

	users, err := l.projectManage.GetProjectUserList(req.ProjectID)
	if err != nil {
		return nil, err
	}

	return &types.GetProjectUserListResponse{
		Users: users,
	}, nil
	return
}
