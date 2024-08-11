package file

import (
	"context"
	"github.com/colinrs/protohub/internal/repository/files"
	"github.com/colinrs/protohub/pkg/utils"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	fileRepository files.FileRepository
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		fileRepository: files.NewFileRepository(ctx, svcCtx),
	}
}

func (l *DetailLogic) Detail(req *types.FileDetailRequest) (resp *types.FileDetailResponse, err error) {
	fileDetail, err := l.fileRepository.FindById(uint(req.FileID))
	if err != nil {
		return nil, err
	}
	fileContent, err := l.fileRepository.FindContentById(uint(req.FileID))
	if err != nil {
		return nil, err
	}
	resp = &types.FileDetailResponse{
		ProjectName: fileDetail.ProjectName,
		ServiceName: fileDetail.ServiceName,
		Branch:      fileDetail.Branch,
		FileName:    fileDetail.FileName,
		FileID:      int(fileDetail.FileID),
		Content:     fileContent.FileContent,
		Creator:     fileDetail.Creator,
		UpdateAt:    fileDetail.UpdatedAt.Format(utils.TimeLayout),
	}
	return
}
