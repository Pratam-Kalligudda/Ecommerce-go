package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignup struct {
	Phone string `json:"phone"`
	UserLogin
}
