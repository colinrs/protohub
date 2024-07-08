package publicuser

import (
	"context"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginBySmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginBySmsLogic {
	return &LoginBySmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginBySmsLogic) LoginBySms(req *types.LoginBySmsRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
