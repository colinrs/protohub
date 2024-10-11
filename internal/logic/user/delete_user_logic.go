package user

import (
	"context"

	"github.com/colinrs/protohub/internal/repository"
	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	userRepository repository.UserRepository
	db             *gorm.DB
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		userRepository: repository.NewUserRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUsersByIDsRequest) error {
	return l.userRepository.DeleteUser(l.db, req.Ids)
}
