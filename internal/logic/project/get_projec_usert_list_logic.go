package project

import (
	"context"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjecUsertListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProjecUsertListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjecUsertListLogic {
	return &GetProjecUsertListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjecUsertListLogic) GetProjecUsertList(req *types.GetProjectUserListRequest) (resp *types.GetProjectUserListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
