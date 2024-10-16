package project

import (
	"context"

	"github.com/colinrs/protohub/internal/repository"
	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	db                *gorm.DB
	projectRepository repository.ProjectRepository
}

func NewGetProjectByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectByIdLogic {
	return &GetProjectByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		projectRepository: repository.NewProjectRepository(ctx, svcCtx),
		db:                svcCtx.DB.WithContext(ctx),
	}
}

func (l *GetProjectByIdLogic) GetProjectById(req *types.GetProjectByIDRequest) (resp *types.GetProjectByIDResponse, err error) {

	pro, err := l.projectRepository.FindProjectByID(l.db, uint64(req.ID))
	if err != nil {
		return nil, err
	}

	resp = &types.GetProjectByIDResponse{
		ID:     uint32(pro.ID),
		Name:   pro.ProjectName,
		Remark: pro.Remark,
		Sort:   pro.Sort,
	}
	return resp, nil
}
