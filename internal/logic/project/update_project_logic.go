package project

import (
	"context"

	"github.com/colinrs/protohub/pkg/code"

	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/repository"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	db                *gorm.DB
	projectRepository repository.ProjectRepository
}

func NewUpdateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProjectLogic {
	return &UpdateProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		projectRepository: repository.NewProjectRepository(ctx, svcCtx),
		db:                svcCtx.DB.WithContext(ctx),
	}
}

func (l *UpdateProjectLogic) UpdateProject(req *types.UpdateProjectRequest) error {
	project, err := l.projectRepository.FindProjectByID(l.db, req.ID)
	if err != nil {
		return err
	}
	if project == nil {
		return code.ErrProjectNotExist
	}
	project.ProjectName = req.Name
	project.Remark = req.Remark
	project.Sort = req.Sort
	err = l.projectRepository.UpdateProject(l.db, project)
	if err != nil {
		return err
	}
	return nil
}
