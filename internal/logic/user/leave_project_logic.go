package user

import (
	"context"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LeaveProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLeaveProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveProjectLogic {
	return &LeaveProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LeaveProjectLogic) LeaveProject(req *types.LeaveProjectRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
