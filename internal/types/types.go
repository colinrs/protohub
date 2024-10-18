// Code generated by goctl. DO NOT EDIT.
package types

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,max=30"`
	NewPassword string `json:"new_password" validate:"required,max=30"`
}

type CreateProjectRequest struct {
	Name   string `json:"name" validate:"max=30"`
	Remark string `json:"remark,optional" validate:"omitempty,max=200"`
}

type CreateUserRequest struct {
	Status      uint32 `json:"status,options=1|2" validate:"omitempty,lt=3"`
	Username    string `json:"username" validate:"omitempty,max=50"`
	Password    string `json:"password" validate:"min=6"`
	Description string `json:"description,optional" validate:"omitempty,max=100"`
	Mobile      string `json:"mobile" validate:"max=18"`
	Email       string `json:"email" validate:"email,max=80"`
}

type DeleteProjectRequest struct {
	IDs []uint `json:"ids"`
}

type DeleteUsersByIDsRequest struct {
	Ids []uint64 `json:"ids"`
}

type FileDetailRequest struct {
	FileID uint `form:"file_id"`
}

type FileDetailResponse struct {
	ProjectName string `json:"project_name"`
	ServiceName string `json:"service_name"`
	Branch      string `json:"branch"`
	FileName    string `json:"file_name"`
	FileID      uint   `json:"file_id"`
	Content     string `json:"content"`
	Creator     string `json:"creator"`
	UpdateAt    string `json:"update_at"`
}

type FileDownloadRequest struct {
	FileID int `form:"file_id"`
}

type FileDownloadResponse struct {
}

type FileListData struct {
	ProjectName string `json:"project_name"`
	ServiceName string `json:"service_name"`
	Branch      string `json:"branch"`
	FileName    string `json:"file_name"`
	FileID      int    `json:"file_id"`
	Creator     string `json:"creator"`
	UpdateAt    string `json:"update_at"`
}

type FileListRequest struct {
	Page        int    `form:"page,optional,default=1"`
	PageSize    int    `form:"page_size,optional,default=10"`
	ProjectName string `form:"project_name"`
	ServiceName string `form:"service_name"`
}

type FileListResponse struct {
	List  []*FileListData `json:"list"`
	Total int             `json:"total"`
}

type FileUploadRequest struct {
	ProjectName string `form:"project_name"`
	ServiceName string `form:"service_name"`
	Branch      string `form:"branch"`
	FileName    string `form:"file_name"`
	Creator     string `form:"creator"`
}

type FileUploadResponse struct {
}

type GetProjectByIDRequest struct {
	ID uint `form:"id"`
}

type GetProjectByIDResponse struct {
	ID     uint   `json:"id,omitempty"`
	Status uint32 `json:"status,omitempty"`
	Name   string `json:"name,omitempty"`
	Remark string `json:"remark,omitempty"`
	Sort   uint32 `json:"sort,omitempty"`
}

type GetProjectListData struct {
	ID     uint   `json:"id"`
	Status uint32 `json:"status"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Sort   uint32 `json:"sort"`
}

type GetProjectListRequest struct {
	Page     int `form:"page,optional,default=1"`
	PageSize int `form:"page_size,optional,default=10"`
}

type GetProjectListResponse struct {
	Total int                   `json:"total"`
	List  []*GetProjectListData `json:"list"`
}

type GetProjectUserListData struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	RoleName string `json:"role_name"`
	RoleId   uint   `json:"role_id"`
}

type GetProjectUserListRequest struct {
	Page      int  `form:"page,optional,default=1"`
	PageSize  int  `form:"page_size,optional,default=10"`
	ProjectID uint `form:"project_id"`
}

type GetProjectUserListResponse struct {
	Total int                       `json:"total"`
	List  []*GetProjectUserListData `json:"list"`
}

type GetRoleByCodeRequest struct {
	Code string `form:"code"`
}

type GetRoleByCodeResponse struct {
	ID     uint   `json:"id,omitempty"`
	Status uint32 `json:"status,omitempty"`
	Name   string `json:"name,omitempty"`
	Code   string `json:"code,omitempty"`
	Remark string `json:"remark,omitempty"`
	Sort   uint32 `json:"sort,omitempty"`
}

type GetRoleListData struct {
	ID     uint   `json:"id"`
	Status uint32 `json:"status"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Remark string `json:"remark"`
	Sort   uint32 `json:"sort"`
}

type GetRoleListRequest struct {
	Page     int    `form:"page,optional,default=1"`
	PageSize int    `form:"page_size,optional,default=10"`
	Name     string `form:"name,optional"`
	Code     string `form:"code,optional"`
}

type GetRoleListResponse struct {
	Total int                `json:"total"`
	List  []*GetRoleListData `json:"list"`
}

type GetUserByIDRequest struct {
	Id uint64 `form:"id" validate:"len=36"`
}

type GetUserByIDResponse struct {
	ID          uint     `json:"id"`
	Username    string   `json:"user_name"`
	Avatar      string   `json:"avatar"`
	Description string   `json:"desc"`
	RoleCode    []string `json:"role_code"`
	RoleIds     []uint64 `json:"role_ids"`
	Mobile      string   `json:"mobile"`
	Email       string   `json:"email"`
	Projects    []string `json:"projects"`
	ProjectIDs  []uint64 `json:"project_ids"`
	Status      uint32   `json:"status,optional" validate:"omitempty,lt=20"`
}

type JoinProjectRequest struct {
	UserID    uint `json:"user_id" validate:"required,numeric"`
	ProjectID uint `json:"project_id" validate:"required,numeric"`
}

type LeaveProjectRequest struct {
	UserID    uint `json:"user_id" validate:"required,numeric"`
	ProjectID uint `json:"project_id" validate:"required,numeric"`
}

type LoginInfo struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
	Expire uint64 `json:"expire"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,alphanum,max=20"`
	Password string `json:"password" validate:"required,max=30,min=6"`
}

type LoginResponse struct {
	UserId    uint   `json:"user_id"`
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expire_at"`
}

type RefreshTokenResponse struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expire_at"`
}

type RoleInfoSimple struct {
	RoleName string `json:"roleName"`
	Value    string `json:"value"`
}

type UpdateProjectRequest struct {
	ID     uint   `json:"id"`
	Status uint32 `json:"status,omitempty" validate:"required,oneof=1 2"`
	Name   string `json:"name" validate:"max=30"`
	Remark string `json:"remark,optional" validate:"omitempty,max=200"`
	Sort   uint32 `json:"sort,optional"`
}

type UpdateUserProfileRequest struct {
	Avatar string `json:"avatar" validate:"omitempty,max=300"`
	Mobile string `json:"mobile" validate:"omitempty,numeric,max=18"`
	Email  string `json:"email" validate:"omitempty,email,max=100"`
}

type UpdateUserRequest struct {
	UserID      uint   `json:"user_id" validate:"omitempty,min=1"`
	Status      uint32 `json:"status,optional" validate:"omitempty,lt=20"`
	Username    string `json:"username,optional" validate:"omitempty,max=50"`
	Description string `json:"description,optional" validate:"omitempty,max=100"`
	Mobile      string `json:"mobile,optional" validate:"max=18"`
	Email       string `json:"email,optional" validate:"email,max=80"`
}

type UserListData struct {
	ID          uint   `json:"id"`
	Status      uint32 `json:"status,optional" validate:"omitempty,lt=20"`
	Username    string `json:"username" validate:"omitempty,max=50"`
	Description string `json:"description,optional" validate:"omitempty,max=100"`
	Mobile      string `json:"mobile" validate:"max=18"`
	Email       string `json:"email" validate:"email,max=80"`
}

type UserListRequest struct {
	Page     int    `form:"page,optional,default=1"`
	PageSize int    `form:"page_size,optional,default=10"`
	Username string `form:"username,optional" validate:"omitempty,alphanum,max=20"`
	Mobile   string `form:"mobile,optional" validate:"omitempty,numeric,max=18"`
	Email    string `form:"email,optional" validate:"omitempty,email,max=100"`
}

type UserListResponse struct {
	Total int             `json:"total"`
	List  []*UserListData `json:"list"`
}
