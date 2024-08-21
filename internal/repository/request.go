package repository

import (
	"github.com/colinrs/protohub/internal/models"
)

type CreateRoleRequest struct {
	Name   string
	Remark string
}

type ListRoleResponse struct {
	Total int
	List  []*models.Role
}

type ListUserResponse struct {
	Total int
	List  []*models.UserTableModel
}

type UserLoginResponse struct {
	Token     string
	ExpiredAt int64
}
