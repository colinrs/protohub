package repository

import (
	"context"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type FileRepository interface {
	QueryFileList(query *FileListQuery, offset, limit int) ([]*models.PbFileTableModel, error)

	Create(pbFilePO *models.PbFileTableModel) error
	CreateFileWithContent(pbFilePO *models.PbFileTableModel, content []byte) error
	Update(id uint, pbFilePO *models.PbFileTableModel) error
	UpdateFileContent(filePO *models.FileContentTableModel) error
	Delete(id uint) error
	FindById(id uint) (*models.PbFileTableModel, error)
	FindContentById(id uint) (*models.FileContentTableModel, error)
	List(offset, limit int) ([]*models.PbFileTableModel, error)
	Count(query *FileListQuery) (int64, error)
}

// pbFileTableRepositoryImpl 实现仓库接口
type pbFileTableRepositoryImpl struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	db *gorm.DB
}

func NewFileRepository(ctx context.Context, svcCtx *svc.ServiceContext) FileRepository {
	return &pbFileTableRepositoryImpl{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
		svcCtx: svcCtx,
		db:     svcCtx.DB.WithContext(ctx),
	}
}

type FileListQuery struct {
	ProjectName string
	ServiceName string
	Branch      string
}

func (repo *pbFileTableRepositoryImpl) QueryFileList(query *FileListQuery, offset, limit int) (
	[]*models.PbFileTableModel, error) {
	fileList := make([]*models.PbFileTableModel, 0)
	err := repo.db.Where(repo.buildFileListQuery(query)).
		Offset(offset).Limit(limit).Find(&fileList).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fileList, nil
		}
		return nil, err
	}
	return fileList, nil
}

func (repo *pbFileTableRepositoryImpl) buildFileListQuery(query *FileListQuery) map[string]interface{} {
	where := map[string]interface{}{}
	if query.ProjectName != "" {
		where[models.ProjectNameKey] = query.ProjectName
	}
	if query.ServiceName != "" {
		where[models.ServiceNameKey] = query.ServiceName
	}
	if query.Branch != "" {
		where[models.BranchKey] = query.Branch
	}
	return where
}

// Create 创建记录
func (repo *pbFileTableRepositoryImpl) Create(data *models.PbFileTableModel) error {
	return repo.db.Save(data).Error
}

func (repo *pbFileTableRepositoryImpl) CreateFileWithContent(pbFilePO *models.PbFileTableModel, content []byte) error {
	fileContentPO := models.FileContentTableModel{
		FileName:    pbFilePO.FileName,
		FileContent: string(content),
		Creator:     pbFilePO.Creator,
	}
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&fileContentPO).Error; err != nil {
			repo.Logger.Errorf("create file content err", err.Error())
			return err
		}
		pbFilePO.FileID = fileContentPO.ID
		if err := tx.Save(pbFilePO).Error; err != nil {
			repo.Logger.Errorf("create file meta err", err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		repo.Logger.Errorf("create file err", err.Error())
		return err
	}
	return nil

}

// Update 更新记录
func (repo *pbFileTableRepositoryImpl) Update(id uint, data *models.PbFileTableModel) error {
	var record *models.PbFileTableModel
	if err := repo.db.Where("id = ?", id).First(&record).Error; err != nil {
		return err
	}
	data.ID = id
	return repo.db.Save(&data).Error
}

// UpdateFileContent 更新记录
func (repo *pbFileTableRepositoryImpl) UpdateFileContent(data *models.FileContentTableModel) error {
	if err := repo.db.Save(data).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除记录
func (repo *pbFileTableRepositoryImpl) Delete(id uint) error {
	var record *models.PbFileTableModel
	if err := repo.db.Where("id = ?", id).Delete(&record).Error; err != nil {
		return err
	}
	return nil
}

// FindById 通过ID查找记录
func (repo *pbFileTableRepositoryImpl) FindById(id uint) (*models.PbFileTableModel, error) {
	var record *models.PbFileTableModel
	if err := repo.db.Where("file_id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (repo *pbFileTableRepositoryImpl) FindContentById(id uint) (*models.FileContentTableModel, error) {
	var record *models.FileContentTableModel
	if err := repo.db.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

// List ...
func (repo *pbFileTableRepositoryImpl) List(offset, limit int) ([]*models.PbFileTableModel, error) {
	var records []*models.PbFileTableModel
	if err := repo.db.Offset(offset).Limit(limit).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (repo *pbFileTableRepositoryImpl) Count(query *FileListQuery) (int64, error) {
	var count int64
	if err := repo.db.Model(models.PbFileTableModel{}).Where(repo.buildFileListQuery(query)).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
