package publicuser

import (
	"context"
	"fmt"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/repository"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/colinrs/protohub/pkg/code"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterByEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	db             *gorm.DB
	userRepository repository.UserRepository
}

func NewRegisterByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterByEmailLogic {
	return &RegisterByEmailLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		userRepository: repository.NewUserRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *RegisterByEmailLogic) RegisterByEmail(req *types.RegisterByEmailRequest) error {
	key := fmt.Sprintf("registerByEmail:%s", req.Email)
	value, err := l.svcCtx.RedisClient.GetCtx(l.ctx, key)
	if err != nil {
		return err
	}
	if value != req.Captcha {
		return code.ErrVerifyCode
	}
	user, err := l.userRepository.FindUser(l.db, &models.UserTableModel{
		Email: req.Email}, 0, 1)
	if err != nil {
		return err
	}
	if len(user.List) > 0 {
		return code.ErrDuplicateUser
	}
	// Assuming passwd is the password string to be encrypted
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := &models.UserTableModel{
		UserName: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	_, err = l.userRepository.CreateUser(l.db, newUser)
	if err != nil {
		return err
	}
	return nil
}
