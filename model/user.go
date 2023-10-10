package model

type User struct {
	ID       string
	Email    string
	Password string
}

type UserCreateRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAuthResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
