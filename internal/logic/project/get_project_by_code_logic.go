package project

import (
	"context"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectByCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProjectByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectByCodeLogic {
	return &GetProjectByCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectByCodeLogic) GetProjectByCode(req *types.GetProjectByIDRequest) (resp *types.GetProjectByIDResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
