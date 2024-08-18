package repository

import (
	"context"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"

	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(db *gorm.DB, req *models.Role) (*models.Role, error)
	FindRole(db *gorm.DB, req *models.Role) (*models.Role, error)
}

type roleRepositoryImpl struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleRepository(ctx context.Context, svcCtx *svc.ServiceContext) RoleRepository {
	return &roleRepositoryImpl{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
		svcCtx: svcCtx,
	}
}

func (r *roleRepositoryImpl) CreateRole(db *gorm.DB, role *models.Role) (*models.Role, error) {
	err := db.Create(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepositoryImpl) FindRole(db *gorm.DB, req *models.Role) (*models.Role, error) {
	return nil, nil
}

func (r *roleRepositoryImpl) FindRoleByCode(db *gorm.DB, req *models.Role) (*models.Role, error) {
	return nil, nil
}
