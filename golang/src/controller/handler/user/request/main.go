package request

type AddUsersReq struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,gte=8"`
}

type UpdateUserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,gte=8"`
	Icon     string `json:"icon" validate:"required"`
}
