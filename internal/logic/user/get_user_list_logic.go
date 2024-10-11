package user

import (
	"context"

	"github.com/colinrs/protohub/pkg/utils"

	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/repository"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	userRepository repository.UserRepository
	db             *gorm.DB
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		userRepository: repository.NewUserRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListRequest) (resp *types.UserListResponse, err error) {
	offset, limit := utils.PageToOffsetLimit(req.Page, req.PageSize)
	userList, err := l.userRepository.FindUser(l.db, &models.UserTableModel{
		UserName: req.Username,
		Email:    req.Email,
		Mobile:   req.Mobile,
	}, offset, limit)
	if err != nil {
		return nil, err
	}
	resp = &types.UserListResponse{
		Total: userList.Total,
	}
	resp.List = make([]*types.UserListData, 0, len(userList.List))
	for _, user := range userList.List {
		resp.List = append(resp.List, &types.UserListData{
			ID:       uint64(user.ID),
			Username: user.UserName,
			Email:    user.Email,
			Mobile:   user.Mobile,
			Status:   user.UserStatus,
		})
	}
	return
}
