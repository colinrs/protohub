package project

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
