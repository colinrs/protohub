package user

import (
	"context"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordByEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordByEmailLogic {
	return &ResetPasswordByEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordByEmailLogic) ResetPasswordByEmail(req *types.ResetPasswordByEmailRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
