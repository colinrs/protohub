package project

import (
	"context"

	"github.com/colinrs/protohub/pkg/utils"

	"github.com/colinrs/protohub/internal/repository"
	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	db                *gorm.DB
	projectRepository repository.ProjectRepository
}

func NewGetProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectListLogic {
	return &GetProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		projectRepository: repository.NewProjectRepository(ctx, svcCtx),
		db:                svcCtx.DB.WithContext(ctx),
	}
}

func (l *GetProjectListLogic) GetProjectList(req *types.GetProjectListRequest) (resp *types.GetProjectListResponse, err error) {
	offset, limit := utils.PageToOffsetLimit(req.Page, req.PageSize)
	projectList, err := l.projectRepository.FindProjectList(l.db, nil, offset, limit)
	if err != nil {
		return nil, err
	}
	resp = &types.GetProjectListResponse{
		List:  make([]*types.GetProjectListData, 0, len(projectList.List)),
		Total: projectList.Total,
	}
	for _, v := range projectList.List {
		resp.List = append(resp.List, &types.GetProjectListData{
			ID:     v.ID,
			Name:   v.ProjectName,
			Remark: v.Remark,
			Sort:   v.Sort,
		})
	}
	return resp, nil
}
