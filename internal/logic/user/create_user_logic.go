package user

import (
	"context"

	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/repository"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx            context.Context
	svcCtx         *svc.ServiceContext
	userRepository repository.RoleRepository
	db             *gorm.DB
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		userRepository: repository.NewRoleRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
