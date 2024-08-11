package role

import (
	"context"
	"github.com/colinrs/protohub/internal/repository/roles"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	roleRepository roles.RoleRepository
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		roleRepository: roles.NewRoleRepository(ctx, svcCtx),
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.CreateRoleRequest) error {

	return nil
}
