package manage

import (
	"context"

	"github.com/colinrs/protohub/internal/types"

	"github.com/colinrs/protohub/internal/svc"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/repository"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ProjectUserList struct {
	Name     string
	UserList []*models.UserTableModel
	Total    int
}

type ProjectManage interface {
	GetProjectUserList(projectID uint) ([]*types.GetProjectUserListData, error)
}

type projectManageImpl struct {
	db                *gorm.DB
	ctx               context.Context
	logger            logx.Logger
	userRepository    repository.UserRepository
	projectRepository repository.ProjectRepository
	roleRepository    repository.RoleRepository
}

func NewProjectManage(ctx context.Context, svcCtx *svc.ServiceContext) ProjectManage {
	return &projectManageImpl{
		ctx:               ctx,
		logger:            logx.WithContext(ctx),
		userRepository:    repository.NewUserRepository(ctx, svcCtx),
		projectRepository: repository.NewProjectRepository(ctx, svcCtx),
		roleRepository:    repository.NewRoleRepository(ctx, svcCtx),
		db:                svcCtx.DB.WithContext(ctx),
	}
}

func (l *projectManageImpl) GetProjectUserList(projectID uint) ([]*types.GetProjectUserListData, error) {

	req := &models.UserProjectTableModel{
		ProjectID: projectID,
	}
	userIDList, err := l.projectRepository.ProjectUserList(l.db, req)
	if err != nil {
		return nil, err
	}

	if len(userIDList) == 0 {
		return nil, nil
	}

	users, err := l.userRepository.FindUserByID(l.db, userIDList)
	if err != nil {
		return nil, err
	}
	userIDList = make([]uint, 0, len(users))
	for _, user := range users {
		userIDList = append(userIDList, user.ID)
	}

	if len(userIDList) == 0 {
		return nil, nil
	}
	userMap := make(map[uint]*models.UserTableModel)
	for _, user := range users {
		userMap[user.ID] = user
	}
	userIDsRols, err := l.roleRepository.FindUserRoleList(l.db, userIDList, projectID)
	if err != nil {
		return nil, err
	}
	resp := make([]*types.GetProjectUserListData, 0, len(userIDList))
	for _, user := range userIDsRols {
		resp = append(resp, &types.GetProjectUserListData{
			ID:       uint32(user.UserID),
			Name:     userMap[user.UserID].UserName,
			RoleName: user.RoleCode,
			RoleId:   uint32(user.RoleId),
		})
	}
	return users, nil
}

func (l *projectManageImpl) ProjectUserList(db *gorm.DB, req *models.UserProjectTableModel) ([]uint, error) {
	return l.projectRepository.ProjectUserList(db, req)
}
