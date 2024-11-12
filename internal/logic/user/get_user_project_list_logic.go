package user

import (
	"context"

	"github.com/colinrs/protohub/pkg/utils"

	"github.com/colinrs/protohub/internal/manage"

	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	db            *gorm.DB
	projectManage manage.ProjectManage
}

func NewGetUserProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProjectListLogic {
	return &GetUserProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		projectManage: manage.NewProjectManage(ctx, svcCtx),
		db:            svcCtx.DB.WithContext(ctx),
	}
}

func (l *GetUserProjectListLogic) GetUserProjectList(req *types.UserProjectListRequest) (resp *types.UserProjectListResponse, err error) {
	offset, limit := utils.PageToOffsetLimit(req.Page, req.PageSize)
	userProjectListData, total, err := l.projectManage.GetUserProject(req.UserID, offset, limit)
	if err != nil {
		return nil, err
	}
	resp = &types.UserProjectListResponse{
		Total: total,
		List:  make([]*types.UserProjectListData, 0, len(userProjectListData)),
	}

	for _, userProject := range userProjectListData {
		resp.List = append(resp.List, &types.UserProjectListData{
			ID:     userProject.ID,
			Name:   userProject.Name,
			Status: userProject.Status,
			Remark: userProject.Remark,
		})
	}
	return resp, nil
}
