package role

import (
	"context"

	"github.com/colinrs/protohub/pkg/utils"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/repository"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	db     *gorm.DB

	roleRepository repository.RoleRepository
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		roleRepository: repository.NewRoleRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.CreateRoleRequest) error {
	roleCode, err := utils.GenerateRandomString(12)
	if err != nil {
		return err
	}
	role := &models.Role{
		Name:       req.Name,
		RoleStatus: models.StatusNormal,
		Code:       roleCode,
	}
	_, err = l.roleRepository.CreateRole(l.db, role)
	if err != nil {
		return err
	}
	return nil
}
