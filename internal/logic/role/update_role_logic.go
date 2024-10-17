package role

import (
	"context"

	"github.com/colinrs/protohub/internal/models"

	"github.com/colinrs/protohub/internal/repository"
	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	db             *gorm.DB
	roleRepository repository.RoleRepository
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		roleRepository: repository.NewRoleRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleRequest) error {
	return l.roleRepository.UpdateRole(l.db, &models.Role{
		Model: gorm.Model{
			ID: req.ID,
		},
		Name:       req.Name,
		Remark:     req.Remark,
		Sort:       req.Sort,
		RoleStatus: uint8(req.Status), // req.Status,
	})
}
