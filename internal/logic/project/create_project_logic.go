package project

import (
	"context"

	"github.com/colinrs/protohub/pkg/code"

	"github.com/colinrs/protohub/internal/models"

	"github.com/colinrs/protohub/internal/repository"
	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	db *gorm.DB

	projectRepository repository.ProjectRepository
}

func NewCreateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProjectLogic {
	return &CreateProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		projectRepository: repository.NewProjectRepository(ctx, svcCtx),
		db:                svcCtx.DB.WithContext(ctx),
	}
}

func (l *CreateProjectLogic) CreateProject(req *types.CreateProjectRequest) error {
	proj, err := l.projectRepository.FindProjectByQuery(l.db, &models.Project{ProjectName: req.Name})
	if err != nil {
		return err
	}
	if proj != nil {
		return code.ErrDuplicateProject
	}
	_, err = l.projectRepository.CreateProject(l.db, &models.Project{
		ProjectName:   req.Name,
		Remark:        req.Remark,
		ProjectStatus: models.ProjectStatusNormal,
	})
	if err != nil {
		return err
	}
	return nil
}
