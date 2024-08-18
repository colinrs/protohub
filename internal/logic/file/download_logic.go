package file

import (
	"context"
	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/repository"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx            context.Context
	svcCtx         *svc.ServiceContext
	fileRepository repository.FileRepository
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		fileRepository: repository.NewFileRepository(ctx, svcCtx),
	}
}

func (l *DownloadLogic) Download(req *types.FileDownloadRequest) (*models.FileContentTableModel, error) {
	file, err := l.fileRepository.FindContentById(uint(req.FileID))
	if err != nil {
		return nil, err
	}
	return file, nil
}
