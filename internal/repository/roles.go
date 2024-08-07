package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type RoleRepository interface {
	CreateRole(req *CreateRoleRequest) error
	FindById(id uint) (*Role, error)
	FindByName(name string) (*Role, error)
	List(offset, limit int) ([]*Role, error)
	Count() (int64, error)
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
