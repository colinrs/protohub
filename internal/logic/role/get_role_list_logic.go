package role

import (
	"context"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/pkg/utils"

	"gorm.io/gorm"

	"github.com/colinrs/protohub/internal/repository"

	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	db *gorm.DB

	roleRepository repository.RoleRepository
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		roleRepository: repository.NewRoleRepository(ctx, svcCtx),
		db:             svcCtx.DB.WithContext(ctx),
	}
}

func (l *GetRoleListLogic) GetRoleList(req *types.GetRoleListRequest) (resp *types.GetRoleListResponse, err error) {
	offset, limit := utils.PageToOffsetLimit(req.Page, req.PageSize)
	roles, err := l.roleRepository.FindRole(l.db, &models.Role{
		Name: req.Name,
		Code: req.Code,
	}, offset, limit)
	if err != nil {
		return nil, err
	}
	var list []*types.GetRoleListData
	for _, role := range roles.List {
		list = append(list, &types.GetRoleListData{
			ID:     uint32(role.ID),
			Status: uint32(role.RoleStatus),
			Name:   role.Name,
			Code:   role.Code,
			Remark: role.Remark,
			Sort:   role.Sort,
		})
	}
	return &types.GetRoleListResponse{Total: roles.Total, List: list}, nil
}
