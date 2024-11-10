package repository

import (
	"context"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(db *gorm.DB, req *models.UserTableModel) (*models.UserTableModel, error)
	FindUserByID(db *gorm.DB, id []uint) ([]*models.UserTableModel, error)
	FindUser(db *gorm.DB, req *models.UserTableModel, offset, limit int) (*ListUserResponse, error)
	DeleteUser(db *gorm.DB, ids []uint64) error
	UpdateUser(db *gorm.DB, req *models.UserTableModel) error
	UserJoinProject(db *gorm.DB, req *models.UserProjectTableModel) error
	UserProjectList(db *gorm.DB, userID uint, offset, limit int) ([]*models.UserProjectTableModel, int, error)
	UserLeaveProject(db *gorm.DB, req *models.UserProjectTableModel) error
	UserJoinRole(db *gorm.DB, req *models.UserRolesTableModel) error
	UserLeaveRole(db *gorm.DB, req *models.UserRolesTableModel) error
	UserLogin(db *gorm.DB, req *models.UserTableModel) (*models.UserTableModel, error)
}

type userRepositoryImpl struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepository(ctx context.Context, svcCtx *svc.ServiceContext) UserRepository {
	return &userRepositoryImpl{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
		svcCtx: svcCtx,
	}
}

func (r *userRepositoryImpl) CreateUser(db *gorm.DB, user *models.UserTableModel) (*models.UserTableModel, error) {
	err := db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) FindUser(db *gorm.DB, req *models.UserTableModel, offset, limit int) (*ListUserResponse, error) {
	resp := &ListUserResponse{}

	err := db.Model(&models.UserTableModel{}).Where(req).Offset(offset).Limit(limit).Find(&resp.List).Error
	if err != nil {
		return nil, err
	}
	var count int64
	err = db.Model(&models.UserTableModel{}).Where(req).Count(&count).Error
	if err != nil {
		return nil, err
	}
	resp.Total = int(count)
	return resp, nil
}

func (r *userRepositoryImpl) DeleteUser(db *gorm.DB, ids []uint64) error {
	return db.Delete(&models.UserTableModel{}, ids).Error
}

func (r *userRepositoryImpl) UpdateUser(db *gorm.DB, req *models.UserTableModel) error {
	return db.Model(&models.UserTableModel{}).Where("id = ?", req.ID).Updates(req).Error
}

func (r *userRepositoryImpl) UserJoinProject(db *gorm.DB, req *models.UserProjectTableModel) error {
	return db.FirstOrCreate(req).Error
}

func (r *userRepositoryImpl) UserLeaveProject(db *gorm.DB, req *models.UserProjectTableModel) error {
	return db.Delete(&models.UserProjectTableModel{}, req).Error
}

func (r *userRepositoryImpl) UserJoinRole(db *gorm.DB, req *models.UserRolesTableModel) error {
	return db.Create(req).Error
}

func (r *userRepositoryImpl) UserLeaveRole(db *gorm.DB, req *models.UserRolesTableModel) error {
	return db.Delete(&models.UserRolesTableModel{}, req).Error
}

func (r *userRepositoryImpl) UserLogin(db *gorm.DB, req *models.UserTableModel) (*models.UserTableModel, error) {
	var user *models.UserTableModel
	err := db.Where(req).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) FindUserByID(db *gorm.DB, id []uint) ([]*models.UserTableModel, error) {
	var resp []*models.UserTableModel
	err := db.Model(&models.UserTableModel{}).Where("id in ?", id).Find(&resp).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *userRepositoryImpl) UserProjectList(db *gorm.DB, userID uint, offset, limit int) ([]*models.UserProjectTableModel, int, error) {
	var resp []*models.UserProjectTableModel
	var totol int64

	err := db.Model(&models.UserProjectTableModel{}).Where("user_id = ?", userID).Count(&totol).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Model(&models.UserProjectTableModel{}).Where("user_id = ?", userID).
		Offset(offset).Limit(limit).Find(&resp).Error
	if err != nil {
		return nil, 0, err
	}
	return resp, int(totol), nil
}
