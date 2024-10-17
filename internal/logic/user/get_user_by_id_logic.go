package user

import (
	"context"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/repository"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/colinrs/protohub/pkg/code"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	userRepository repository.UserRepository
	db             *gorm.DB
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		userRepository: repository.NewUserRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.GetUserByIDRequest) (resp *types.GetUserByIDResponse, err error) {
	userList, err := l.userRepository.FindUser(l.db, &models.UserTableModel{
		Model: gorm.Model{ID: uint(req.Id)},
	}, 0, 1)

	if err != nil {
		return nil, err
	}
	if userList.Total == 0 {
		return nil, code.ErrUserNotFound
	}
	resp = &types.GetUserByIDResponse{
		ID:       userList.List[0].ID,
		Username: userList.List[0].UserName,
		Email:    userList.List[0].Email,
		Status:   userList.List[0].UserStatus,
	}
	return
}
