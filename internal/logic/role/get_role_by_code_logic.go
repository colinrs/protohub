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

type GetRoleByCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	db             *gorm.DB
	roleRepository repository.RoleRepository
}

func NewGetRoleByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByCodeLogic {
	return &GetRoleByCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		roleRepository: repository.NewRoleRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *GetRoleByCodeLogic) GetRoleByCode(req *types.GetRoleByCodeRequest) (resp *types.GetRoleByCodeResponse, err error) {
	roles, err := l.roleRepository.FindRole(l.db, &models.Role{
		Code: req.Code,
	}, 0, 1)
	if err != nil {
		return nil, err
	}
	if roles.Total == 0 || len(roles.List) == 0 {
		return nil, nil
	}
	return &types.GetRoleByCodeResponse{
		ID:     uint32(roles.List[0].ID),
		Status: uint32(roles.List[0].RoleStatus),
		Name:   roles.List[0].Name,
		Code:   roles.List[0].Code,
		Remark: roles.List[0].Remark,
		Sort:   roles.List[0].Sort,
	}, nil
}
