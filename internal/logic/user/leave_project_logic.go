package user

import (
	"context"

	"github.com/colinrs/protohub/internal/manage"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LeaveProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	projectManage manage.ProjectManage
}

func NewLeaveProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveProjectLogic {
	return &LeaveProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		projectManage: manage.NewProjectManage(ctx, svcCtx),
	}
}

func (l *LeaveProjectLogic) LeaveProject(req *types.LeaveProjectRequest) error {
	return l.projectManage.DeleteUserFromProject(req.ProjectID, req.UserID)
}
