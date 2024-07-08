// Code generated by goctl. DO NOT EDIT.
package types

type AccessTokenResponse struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}

type BaseIDInfo struct {
	Id        *uint64 `json:"id,optional"`
	CreatedAt *int64  `json:"createdAt,optional"`
	UpdatedAt *int64  `json:"updatedAt,optional"`
}

type BaseUUIDInfo struct {
	Id        *string `json:"id,optional"`
	CreatedAt *int64  `json:"createdAt,optional"`
	UpdatedAt *int64  `json:"updatedAt,optional"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,max=30"`
	NewPassword string `json:"new_password" validate:"required,max=30"`
}

type CreateRoleRequest struct {
	RoleInfo
}

type CreateUserRequest struct {
	UserInfo
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

type GetPermCodeResponse struct {
	List []string `json:"list"`
}

type GetRoleByIdRequest struct {
	ID uint64 `form:"id"`
}

type GetRoleByIdResponse struct {
	RoleInfo
}

type GetRoleListRequest struct {
	PageLimit
	Name string `form:"name,optional"`
}

type GetRoleListResponse struct {
	Total int         `json:"total"`
	List  []*RoleInfo `json:"list"`
}

type GetUserByIDRequest struct {
	Id string `form:"id" validate:"len=36"`
}

type GetUserByIDResponse struct {
	UserInfo
}

type GetUserInfoResponse struct {
	UUID           *string  `json:"user_id"`
	Username       *string  `json:"username"`
	Nickname       *string  `json:"nickname"`
	Avatar         *string  `json:"avatar"`
	HomePath       *string  `json:"home_path"`
	Description    *string  `json:"desc"`
	RoleName       []string `json:"role_name"`
	DepartmentName string   `json:"department_name,optional"`
}

type GetUserProfileResponse struct {
	Nickname *string `json:"nickname" validate:"omitempty,alphanumunicode,max=10"`
	Avatar   *string `json:"avatar" validate:"omitempty,max=300"`
	Mobile   *string `json:"mobile" validate:"omitempty,numeric,max=18"`
	Email    *string `json:"email" validate:"omitempty,email,max=100"`
}

type IDsReq struct {
	IDs []uint64 `json:"ids"`
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
	Username string `json:"username" validate:"required,alphanum,max=20"`
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

type RoleInfo struct {
	BaseIDInfo
	Trans         string  `json:"trans,optional"`
	Status        *uint32 `json:"status,optional" validate:"omitempty,lt=20"`
	Name          *string `json:"name,optional" validate:"omitempty,max=30"`
	Code          *string `json:"code,optional" validate:"omitempty,max=20"`
	DefaultRouter *string `json:"default_router,optional" validate:"omitempty,max=80"`
	Remark        *string `json:"remark,optional" validate:"omitempty,max=200"`
	Sort          *uint32 `json:"sort,optional" validate:"omitempty,lt=10000"`
}

type RoleInfoResp struct {
	RoleInfo
}

type RoleInfoSimple struct {
	RoleName string `json:"roleName"`
	Value    string `json:"value"`
}

type UpdateRoleRequest struct {
	RoleInfo
}

type UpdateUserProfileRequest struct {
	Nickname *string `json:"nickname" validate:"omitempty,alphanumunicode,max=10"`
	Avatar   *string `json:"avatar" validate:"omitempty,max=300"`
	Mobile   *string `json:"mobile" validate:"omitempty,numeric,max=18"`
	Email    *string `json:"email" validate:"omitempty,email,max=100"`
}

type UpdateUserProfileResponse struct {
	Nickname *string `json:"nickname" validate:"omitempty,alphanumunicode,max=10"`
	Avatar   *string `json:"avatar" validate:"omitempty,max=300"`
	Mobile   *string `json:"mobile" validate:"omitempty,numeric,max=18"`
	Email    *string `json:"email" validate:"omitempty,email,max=100"`
}

type UpdateUserRequest struct {
	UserInfo
}

type UserInfo struct {
	BaseUUIDInfo
	Status       *uint32  `json:"status,optional" validate:"omitempty,lt=20"`
	Username     *string  `json:"username,optional" validate:"omitempty,max=50"`
	Nickname     *string  `json:"nickname,optional" validate:"omitempty,max=40"`
	Password     *string  `json:"password,optional" validate:"omitempty,min=6"`
	Description  *string  `json:"description,optional" validate:"omitempty,max=100"`
	HomePath     *string  `json:"home_path,optional" validate:"omitempty,max=70"`
	RoleIds      []uint64 `json:"role_ids,optional"`
	Mobile       *string  `json:"mobile,optional" validate:"omitempty,max=18"`
	Email        *string  `json:"email,optional" validate:"omitempty,max=80"`
	Avatar       *string  `json:"avatar,optional" validate:"omitempty,max=300"`
	DepartmentId *uint64  `json:"department_id,optional"`
	PositionIds  []uint64 `json:"position_id,optional"`
}

type UserListRequest struct {
	PageLimit
	Username     *string  `json:"username,optional" validate:"omitempty,alphanum,max=20"`
	Nickname     *string  `json:"nickname,optional" validate:"omitempty,alphanumunicode,max=10"`
	Mobile       *string  `json:"mobile,optional" validate:"omitempty,numeric,max=18"`
	Email        *string  `json:"email,optional" validate:"omitempty,email,max=100"`
	RoleIds      []uint64 `json:"role_ids,optional"`
	DepartmentId *uint64  `json:"department_id,optional"`
	PositionId   *uint64  `json:"position_id,optional"`
}

type UserListResponse struct {
	Total int         `json:"total"`
	List  []*UserInfo `json:"list"`
}
