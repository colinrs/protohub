package manage

import (
	"context"

	"github.com/colinrs/protohub/internal/models"
	"github.com/colinrs/protohub/internal/repository"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/internal/types"
	"github.com/colinrs/protohub/pkg/code"

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
	AddUserToProject(projectID, userID uint) error
	DeleteUserFromProject(projectID, userID uint) error
	GetUserProject(userID uint) ([]*types.UserProjectListData, int, error)
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

	userIDList, err := l.projectRepository.ProjectUserList(l.db, projectID)
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
	roleCodes := make([]string, 0, len(userIDsRols))

	for _, user := range userIDsRols {
		roleCodes = append(roleCodes, user.RoleCode)
	}

	roles, err := l.roleRepository.FindRoleByCode(l.db, roleCodes)
	if err != nil {
		return nil, err
	}
	roleMap := make(map[string]*models.Role)
	for _, role := range roles {
		roleMap[role.Code] = role
	}
	resp := make([]*types.GetProjectUserListData, 0, len(userIDList))
	for _, user := range userIDsRols {
		resp = append(resp, &types.GetProjectUserListData{
			ID:       user.UserID,
			Name:     userMap[user.UserID].UserName,
			RoleName: user.RoleCode,
			RoleId:   roleMap[user.RoleCode].ID,
		})
	}
	return resp, nil
}

func (l *projectManageImpl) ProjectUserList(db *gorm.DB, req *models.UserProjectTableModel) ([]uint, error) {
	return l.projectRepository.ProjectUserList(db, req.ProjectID)
}

func (l *projectManageImpl) AddUserToProject(projectID, userID uint) error {
	if projectID == 0 || userID == 0 {
		return code.ErrParam
	}
	userList, err := l.userRepository.FindUserByID(l.db, []uint{userID})
	if err != nil {
		return err
	}
	if len(userList) == 0 {
		return code.ErrUserNotFound
	}
	project, err := l.projectRepository.FindProjectByID(l.db, projectID)
	if err != nil {
		return err
	}
	if project == nil {
		return code.ErrProjectNotExist
	}
	err = l.userRepository.UserJoinProject(l.db, &models.UserProjectTableModel{
		ProjectID: projectID,
		UserID:    userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (l *projectManageImpl) DeleteUserFromProject(projectID, userID uint) error {
	err := l.userRepository.UserLeaveProject(l.db, &models.UserProjectTableModel{
		ProjectID: projectID,
		UserID:    userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (l *projectManageImpl) GetUserProject(userID uint) ([]*types.UserProjectListData, int, error) {
	user, err := l.userRepository.FindUserByID(l.db, []uint{userID})
	if err != nil {
		return nil, 0, err
	}
	if len(user) == 0 {
		return nil, 0, code.ErrUserNotFound
	}
	userProjectTableModel, total, err := l.userRepository.UserProjectList(l.db, userID, 0, 0)
	if err != nil {
		return nil, 0, err
	}
	projectIDs := make([]uint, 0, len(userProjectTableModel))
	for _, userProject := range userProjectTableModel {
		projectIDs = append(projectIDs, userProject.ProjectID)
	}
	projects, err := l.projectRepository.FindProjectsByID(l.db, projectIDs)
	if err != nil {
		return nil, 0, err
	}
	projectMap := make(map[uint]*models.Project)
	for _, project := range projects {
		projectMap[project.ID] = project
	}
	resp := make([]*types.UserProjectListData, 0, len(userProjectTableModel))
	for _, userProject := range userProjectTableModel {
		projectInfo, ok := projectMap[userProject.ProjectID]
		if !ok {
			continue
		}
		resp = append(resp, &types.UserProjectListData{
			ID:     userProject.ProjectID,
			Name:   projectInfo.ProjectName,
			Remark: projectInfo.Remark,
			Status: projectInfo.ProjectStatus,
		})
	}
	return resp, total, nil
}
