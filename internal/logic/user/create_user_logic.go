package user

import (
	"context"

	"github.com/colinrs/protohub/pkg/utils"

	"github.com/colinrs/protohub/pkg/code"

	"github.com/colinrs/protohub/internal/models"

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
	userRepository repository.UserRepository
	db             *gorm.DB
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		userRepository: repository.NewUserRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserRequest) error {

	userList, err := l.userRepository.FindUser(l.db, &models.UserTableModel{
		Email: req.Email}, 0, 1)
	if err != nil {
		return err
	}
	if userList.Total > 0 {
		return code.ErrUserExistBefor
	}
	_, err = l.userRepository.CreateUser(l.db, &models.UserTableModel{
		UserName:    req.Username,
		Password:    utils.HashPassword(req.Password, l.svcCtx.Config.PasswdSecret),
		Email:       req.Email,
		UserStatus:  req.Status,
		Description: req.Description,
		Mobile:      req.Mobile,
	})
	if err != nil {
		return err
	}
	return nil
}
