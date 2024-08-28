package publicuser

import (
	"context"

	"github.com/colinrs/protohub/internal/repository"
	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	db             *gorm.DB
	userRepository repository.UserRepository
}

func NewLoginByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByEmailLogic {
	return &LoginByEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		userRepository: repository.NewUserRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *LoginByEmailLogic) LoginByEmail(req *types.LoginByEmailRequest) (resp *types.LoginResponse, err error) {

	return
}
