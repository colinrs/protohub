package repository

import (
	"context"
	"errors"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(db *gorm.DB, req *models.Role) (*models.Role, error)
	FindRole(db *gorm.DB, req *models.Role, offset, limit int) (*ListRoleResponse, error)
	DeleteRole(db *gorm.DB, ids []uint64) error
	UpdateRole(db *gorm.DB, req *models.Role) error
	FindUserRoleList(db *gorm.DB, userIDs []uint, projectID uint) ([]*models.UserRolesTableModel, error)
	FindRoleByQuery(db *gorm.DB, req *models.Role) (*models.Role, error)
	FindRoleByCode(db *gorm.DB, codes []string) ([]*models.Role, error)
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

func (r *roleRepositoryImpl) FindRole(db *gorm.DB, req *models.Role, offset, limit int) (*ListRoleResponse, error) {
	resp := &ListRoleResponse{}

	err := db.Model(&models.Role{}).Where(req).Offset(offset).Limit(limit).Find(&resp.List).Error
	if err != nil {
		return nil, err
	}
	var count int64
	err = db.Model(&models.Role{}).Where(req).Count(&count).Error
	if err != nil {
		return nil, err
	}
	resp.Total = int(count)
	return resp, nil
}

func (r *roleRepositoryImpl) DeleteRole(db *gorm.DB, ids []uint64) error {
	return db.Delete(&models.Role{}, ids).Error
}

func (r *roleRepositoryImpl) UpdateRole(db *gorm.DB, req *models.Role) error {
	return db.Model(&models.Role{}).Where("id = ?", req.ID).Updates(req).Error
}

func (r *roleRepositoryImpl) FindUserRoleList(db *gorm.DB, userIDs []uint, projectID uint) ([]*models.UserRolesTableModel, error) {
	var list []*models.UserRolesTableModel
	err := db.Model(&models.UserRolesTableModel{}).Where("project_id = ? and user_id in ?", projectID, userIDs).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *roleRepositoryImpl) UserJoinRole(db *gorm.DB, req *models.UserRolesTableModel) error {
	return db.Create(req).Error
}

func (r *roleRepositoryImpl) UserLeaveRole(db *gorm.DB, req *models.UserRolesTableModel) error {
	return db.Delete(&models.UserRolesTableModel{}, req).Error
}

func (r *roleRepositoryImpl) FindRoleByQuery(db *gorm.DB, req *models.Role) (*models.Role, error) {
	err := db.Model(&models.Role{}).Where(req).First(req).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return req, nil
}

func (r *roleRepositoryImpl) FindRoleByCode(db *gorm.DB, codes []string) ([]*models.Role, error) {
	var list []*models.Role
	err := db.Model(&models.Role{}).Where("code in ?", codes).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
