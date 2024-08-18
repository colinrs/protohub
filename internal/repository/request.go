package repository

import "github.com/colinrs/protohub/internal/models"

type CreateRoleRequest struct {
	Name   string
	Remark string
}

type CreateRoleResponse struct {
}

type UpdateRoleRequest struct {
}

type UpdateRoleResponse struct {
}

type DeleteRoleRequest struct {
}

type DeleteRoleResponse struct {
}

type ListRoleRequest struct {
}

type ListRoleResponse struct {
	Total int
	List  []*models.Role
}

type CountRoleRequest struct {
}

type CountRoleResponse struct {
}

type FindRoleRequest struct {
}

type FindRoleResponse struct {
}

type FindByIdRoleRequest struct {
}

type FindByIdRoleResponse struct {
}
