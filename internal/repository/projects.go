package repository

import (
	"context"
	"errors"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	CreateProject(db *gorm.DB, req *models.Project) (*models.Project, error)
	FindProjectByID(db *gorm.DB, id uint64) (*models.Project, error)
	FindProjectList(db *gorm.DB, req *models.Project, offset, limit int) (*ListProjectResponse, error)
	DeleteProject(db *gorm.DB, ids []uint64) error
	UpdateProject(db *gorm.DB, req *models.Project) error
	FindProjectByQuery(db *gorm.DB, req *models.Project) (*models.Project, error)

	ProjectUserList(db *gorm.DB, req *models.UserProjectTableModel) ([]uint, error)
}

type ProjectRepositoryImpl struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectRepository(ctx context.Context, svcCtx *svc.ServiceContext) ProjectRepository {
	return &ProjectRepositoryImpl{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
		svcCtx: svcCtx,
	}
}

func (r *ProjectRepositoryImpl) CreateProject(db *gorm.DB, Project *models.Project) (*models.Project, error) {
	err := db.Create(Project).Error
	if err != nil {
		return nil, err
	}
	return Project, nil
}

func (r *ProjectRepositoryImpl) FindProjectList(db *gorm.DB, req *models.Project, offset, limit int) (*ListProjectResponse, error) {
	resp := &ListProjectResponse{}

	err := db.Model(&models.Project{}).Where(req).Offset(offset).Limit(limit).Find(&resp.List).Error
	if err != nil {
		return nil, err
	}
	var count int64
	err = db.Model(&models.Project{}).Where(req).Count(&count).Error
	if err != nil {
		return nil, err
	}
	resp.Total = int(count)
	return resp, nil
}

func (r *ProjectRepositoryImpl) FindProjectByQuery(db *gorm.DB, req *models.Project) (*models.Project, error) {
	err := db.Model(&models.Project{}).Where(req).First(req).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return req, nil
}

func (r *ProjectRepositoryImpl) DeleteProject(db *gorm.DB, ids []uint64) error {
	return db.Delete(&models.Project{}, ids).Error
}

func (r *ProjectRepositoryImpl) UpdateProject(db *gorm.DB, req *models.Project) error {
	return db.Model(&models.Project{}).Where("id = ?", req.ID).Updates(req).Error
}

func (r *ProjectRepositoryImpl) FindProjectByID(db *gorm.DB, id uint64) (*models.Project, error) {
	resp := &models.Project{}
	err := db.Model(&models.Project{}).Where("id = ?", id).First(resp).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ProjectRepositoryImpl) ProjectUserList(db *gorm.DB, req *models.UserProjectTableModel) ([]uint, error) {
	var resp []uint
	err := db.Model(&models.UserProjectTableModel{}).Where(req).Pluck("user_id", &resp).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}
