// Code generated by goctl. DO NOT EDIT.
package types

type AccessTokenResponse struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,max=30"`
	NewPassword string `json:"new_password" validate:"required,max=30"`
}

type CreateProjectRequest struct {
	Name   string `json:"name" validate:"max=30"`
	Remark string `json:"remark,optional" validate:"omitempty,max=200"`
}

type CreateRoleRequest struct {
	Name   string `json:"name" validate:"max=30"`
	Remark string `json:"remark,optional" validate:"omitempty,max=200"`
}

type CreateUserRequest struct {
	Status      uint32 `json:"status,optional" validate:"omitempty,lt=3"`
	Username    string `json:"username" validate:"omitempty,max=50"`
	Password    string `json:"password" validate:"min=6"`
	Description string `json:"description,optional" validate:"omitempty,max=100"`
	Mobile      string `json:"mobile" validate:"max=18"`
	Email       string `json:"email" validate:"email,max=80"`
}

type DeleteProjectRequest struct {
	IDsReq
}

type DeleteRoleRequest struct {
	IDsReq
}

type DeleteUsersByIDsRequest struct {
	Ids []string `json:"ids"`
}

type FileDetailRequest struct {
	FileID int `form:"file_id"`
}

type FileDetailResponse struct {
	ProjectName string `json:"project_name"`
	ServiceName string `json:"service_name"`
	Branch      string `json:"branch"`
	FileName    string `json:"file_name"`
	FileID      int    `json:"file_id"`
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
	PageLimit
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
	ID uint32 `form:"id"`
}

type GetProjectByIDResponse struct {
	ID     uint32 `json:"id,omitempty"`
	Status uint32 `json:"status,omitempty"`
	Name   string `json:"name,omitempty"`
	Remark string `json:"remark,omitempty"`
	Sort   uint32 `json:"sort,omitempty"`
}

type GetProjectListData struct {
	ID     uint32 `json:"id"`
	Status uint32 `json:"status"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Sort   uint32 `json:"sort"`
}

type GetProjectListRequest struct {
	PageLimit
	Name string `form:"name,optional"`
}

type GetProjectListResponse struct {
	Total int                   `json:"total"`
	List  []*GetProjectListData `json:"list"`
}

type GetRoleByCodeRequest struct {
	Code string `form:"code"`
}

type GetRoleByCodeResponse struct {
	ID     uint32 `json:"id,omitempty"`
	Status uint32 `json:"status,omitempty"`
	Name   string `json:"name,omitempty"`
	Code   string `json:"code,omitempty"`
	Remark string `json:"remark,omitempty"`
	Sort   uint32 `json:"sort,omitempty"`
}

type GetRoleListData struct {
	ID     uint32 `json:"id"`
	Status uint32 `json:"status"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Remark string `json:"remark"`
	Sort   uint32 `json:"sort"`
}

type GetRoleListRequest struct {
	PageLimit
	Name string `form:"name,optional"`
	Code string `form:"code,optional"`
}

type GetRoleListResponse struct {
	Total int                `json:"total"`
	List  []*GetRoleListData `json:"list"`
}

type GetUserByIDRequest struct {
	Id uint64 `form:"id" validate:"len=36"`
}

type GetUserByIDResponse struct {
	Status      uint32 `json:"status,optional" validate:"omitempty,lt=20"`
	Username    string `json:"username" validate:"omitempty,max=50"`
	Description string `json:"description,optional" validate:"omitempty,max=100"`
	Mobile      string `json:"mobile" validate:"max=18"`
	Email       string `json:"email" validate:"email,max=80"`
}

type GetUserInfoResponse struct {
	UserID      uint64   `json:"user_id"`
	Username    string   `json:"username"`
	Avatar      string   `json:"avatar"`
	Description string   `json:"desc"`
	RoleCode    []string `json:"role_code"`
	Mobile      string   `json:"mobile"`
	Email       string   `json:"email"`
	RoleIds     []string `json:"role_ids"`
}

type IDsReq struct {
	IDs []uint64 `json:"ids"`
}

type JoinProjectRequest struct {
	ProjectID int64 `json:"project_id" validate:"required,numeric"`
}

type LeaveProjectRequest struct {
	ProjectID int64 `json:"project_id" validate:"required,numeric"`
}

type LoginByEmailRequest struct {
	Email   string `json:"email" validate:"required,email,max=100"`
	Captcha string `json:"captcha,optional" validate:"omitempty,len=5"`
}

type LoginBySmsRequest struct {
	PhoneNumber string `json:"phone_number"  validate:"required,numeric,max=20"`
	Captcha     string `json:"captcha,optional" validate:"omitempty,len=5"`
}

type LoginInfo struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
	Expire uint64 `json:"expire"`
}

type LoginRequest struct {
	Username  string `json:"username" validate:"required,alphanum,max=20"`
	Password  string `json:"password" validate:"required,max=30,min=6"`
	CaptchaId string `json:"captcha_id"  validate:"required,len=20"`
	Captcha   string `json:"captcha" validate:"required,len=5"`
}

type LoginResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
	Expire uint64 `json:"expire"`
}

type PageLimit struct {
	Page     int `form:"page,optional,default=1"`
	PageSize int `form:"page_size,optional,default=10"`
}

type RefreshTokenResponse struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}

type RegisterByEmailRequest struct {
	Username string `json:"username" validate:"required,max=20"`
	Password string `json:"password" validate:"required,max=30,min=6"`
	Captcha  string `json:"captcha" validate:"required,len=5"`
	Email    string `json:"email" validate:"required,email,max=100"`
}

type RegisterBySmsRequest struct {
	Username    string `json:"username" validate:"required,alphanum,max=20"`
	Password    string `json:"password" validate:"required,max=30,min=6"`
	Captcha     string `json:"captcha" validate:"required,len=5"`
	PhoneNumber string `json:"phone_number"  validate:"required,numeric,max=20"`
}

type RegisterRequest struct {
	Username  string `json:"user_name" validate:"required,alphanum,max=20"`
	Password  string `json:"password" validate:"required,max=30,min=6"`
	CaptchaId string `json:"captcha_id" validate:"required,len=20"`
	Captcha   string `json:"captcha" validate:"required,len=5"`
	Email     string `json:"email" validate:"required,email,max=100"`
}

type ResetPasswordByEmailRequest struct {
	Email    string `json:"email" validate:"email"`
	Captcha  string `json:"captcha"`
	Password string `json:"password"`
}

type ResetPasswordBySmsRequest struct {
	PhoneNumber string `json:"phone_number"`
	Captcha     string `json:"captcha"`
	Password    string `json:"password"`
}

type RoleInfoSimple struct {
	RoleName string `json:"roleName"`
	Value    string `json:"value"`
}

type UpdateProjectRequest struct {
	ID     uint32 `json:"id"`
	Status uint32 `json:"status,omitempty" validate:"required,oneof=1 2"`
	Name   string `json:"name" validate:"max=30"`
	Remark string `json:"remark,optional" validate:"omitempty,max=200"`
	Sort   uint32 `json:"sort,optional"`
}

type UpdateRoleRequest struct {
	ID     uint32 `json:"id"`
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
	UserID      uint64 `json:"user_id" validate:"omitempty,min=1"`
	Status      uint32 `json:"status,optional" validate:"omitempty,lt=20"`
	Username    string `json:"username,optional" validate:"omitempty,max=50"`
	Description string `json:"description,optional" validate:"omitempty,max=100"`
	Mobile      string `json:"mobile,optional" validate:"max=18"`
	Email       string `json:"email,optional" validate:"email,max=80"`
}

type UserListData struct {
	Status      uint32 `json:"status,optional" validate:"omitempty,lt=20"`
	Username    string `json:"username" validate:"omitempty,max=50"`
	Description string `json:"description,optional" validate:"omitempty,max=100"`
	Mobile      string `json:"mobile" validate:"max=18"`
	Email       string `json:"email" validate:"email,max=80"`
}

type UserListRequest struct {
	PageLimit
	Username string `json:"username,optional" validate:"omitempty,alphanum,max=20"`
	Mobile   string `json:"mobile,optional" validate:"omitempty,numeric,max=18"`
	Email    string `json:"email,optional" validate:"omitempty,email,max=100"`
}

type UserListResponse struct {
	Total int             `json:"total"`
	List  []*UserListData `json:"list"`
}
