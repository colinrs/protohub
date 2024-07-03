package file

import (
	"context"
	"io"
	"net/http"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/repository"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/colinrs/pkgx/utils/md5"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	fileRepository repository.FileRepository
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		fileRepository: repository.NewFileRepository(ctx, svcCtx),
	}
}

func (l *UploadLogic) Upload(req *types.FileUploadRequest, r *http.Request) (resp *types.FileUploadResponse, err error) {
	query := &repository.FileListQuery{
		ProjectName: req.ProjectName,
		ServiceName: req.ServiceName,
		Branch:      req.Branch,
	}

	fileList, err := l.fileRepository.QueryFileList(query, 0, 1)
	if err != nil {
		return nil, err
	}
	fileContent, err := l.getFileContent(r)
	if err != nil {
		return nil, err
	}
	md5Value := md5.Md5Hex(md5.Md5(fileContent))
	if len(fileList) == 0 {
		l.Infof("project: %s, service: %s, branch: %s not exit file record, create it, md5:%s",
			req.ProjectName, req.ServiceName, req.Branch, md5Value)
		pbFilePO := &models.PbFileTableModel{
			ProjectName: req.ProjectName,
			ServiceName: req.ServiceName,
			Branch:      req.Branch,
			Creator:     req.Creator,
			FileName:    req.FileName,
			Sign:        md5Value,
		}
		err = l.fileRepository.CreateFileWithContent(pbFilePO, fileContent)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	if md5Value == fileList[0].Sign {
		l.Infof("project: %s, service: %s, branch: %s file content not change",
			req.ProjectName, req.ServiceName, req.Branch)
		return nil, nil
	}
	l.Infof("project: %s, service: %s, branch: %s file id:%d, md5:%s, new md5:%",
		req.ProjectName, req.ServiceName, req.Branch, fileList[0].ID, fileList[0].Sign, md5Value)
	fileP0 := &models.FileContentTableModel{
		Model: gorm.Model{
			ID: fileList[0].ID,
		},
		FileContent: string(fileContent),
	}
	err = l.fileRepository.UpdateFileContent(fileP0)
	if err != nil {
		return nil, err
	}
	resp = &types.FileUploadResponse{}

	return resp, nil
}

func (l *UploadLogic) getFileContent(r *http.Request) (fileContent []byte, err error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileContent, err = io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}
