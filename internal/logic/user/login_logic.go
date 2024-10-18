package user

import (
	"context"
	"fmt"
	"time"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/repository"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/colinrs/protohub/pkg/utils"

	"github.com/rs/xid"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	userRepository repository.UserRepository
	db             *gorm.DB
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		userRepository: repository.NewUserRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {

	user, err := l.userRepository.UserLogin(l.db, &models.UserTableModel{
		UserName: req.Username,
		Password: utils.HashPassword(req.Password, l.svcCtx.Config.PasswdSecret),
	})
	if err != nil {
		return
	}
	now := time.Now()
	expiredAt := time.Now().Add(time.Duration(l.svcCtx.Config.JwtExpired) * time.Second).Unix()
	tid := fmt.Sprintf("%d|%s|%s|%s|%d", user.ID, user.Email, xid.New().String(), now.Unix())
	token, err := utils.GenerateJWT(map[string]interface{}{
		"user_id": user.ID, "email": user.Email, "tid": tid}, []byte(l.svcCtx.Config.JwtSecret), expiredAt)
	if err != nil {
		return
	}
	resp = &types.LoginResponse{
		UserId:    user.ID,
		Token:     token,
		ExpiredAt: expiredAt,
	}
	return
}
