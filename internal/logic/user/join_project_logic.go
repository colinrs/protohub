package user

import (
	"context"

	"github.com/colinrs/protohub/internal/manage"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	projectManage manage.ProjectManage
}

func NewJoinProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinProjectLogic {
	return &JoinProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		projectManage: manage.NewProjectManage(ctx, svcCtx),
	}
}

func (l *JoinProjectLogic) JoinProject(req *types.JoinProjectRequest) error {
	err := l.projectManage.AddUserToProject(req.ProjectID, req.UserID)
	if err != nil {
		return err
	}
	return nil
}
