package file

import (
	"context"

	"github.com/colinrs/protohub/internal/repository"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/colinrs/protohub/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	fileRepository repository.FileRepository
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		fileRepository: repository.NewFileRepository(ctx, svcCtx),
	}
}

func (l *ListLogic) List(req *types.FileListRequest) (resp *types.FileListResponse, err error) {
	query := &repository.FileListQuery{
		ProjectName: req.ProjectName,
		ServiceName: req.ServiceName,
	}
	offset, limit := utils.Page(req.Page, req.PageSize)
	fileList, err := l.fileRepository.QueryFileList(query, offset, limit)
	if err != nil {
		return nil, err
	}
	count, err := l.fileRepository.Count(query)
	if err != nil {
		return nil, err
	}
	list := make([]*types.FileListData, 0, len(fileList))
	for _, f := range fileList {
		list = append(list, &types.FileListData{
			ProjectName: f.ProjectName,
			ServiceName: f.ServiceName,
			Branch:      f.Branch,
			FileName:    f.FileName,
			FileID:      int(f.FileID),
			Creator:     f.Creator,
			UpdateAt:    f.UpdatedAt.Format(utils.TimeLayout),
		})
	}
	resp = &types.FileListResponse{
		List:  list,
		Total: int(count),
	}
	return resp, nil
}
