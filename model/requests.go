package model

// user

type ReqRegister struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// session

type ReqLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
