package user

import (
	"context"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinProjectLogic {
	return &JoinProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinProjectLogic) JoinProject(req *types.JoinProjectRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
