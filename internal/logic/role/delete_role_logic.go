package role

import (
	"context"

	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/repository"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	db             *gorm.DB
	roleRepository repository.RoleRepository
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		roleRepository: repository.NewRoleRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(req *types.DeleteRoleRequest) error {
	return l.roleRepository.DeleteRole(l.db, req.IDs)
}
