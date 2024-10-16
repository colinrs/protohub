package project

import (
	"context"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProjectListLogic {
	return &GetUserProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserProjectListLogic) GetUserProjectList(req *types.GetProjectListByUserRequest) (resp *types.GetProjectListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
