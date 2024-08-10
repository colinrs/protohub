package repository

import (
	"context"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"

	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(req *CreateRoleRequest) error
}

type roleRepositoryImpl struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	db     *gorm.DB
}

func NewRoleRepository(ctx context.Context, svcCtx *svc.ServiceContext) RoleRepository {
	return &roleRepositoryImpl{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
		svcCtx: svcCtx,
		db:     svcCtx.DB.WithContext(ctx),
	}
}

func (r *roleRepositoryImpl) CreateRole(req *CreateRoleRequest) error {
	return nil
}
