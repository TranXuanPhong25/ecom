package models

type LoginResponse struct {
	Token string   `json:"token"`
	User  UserInfo `json:"userInfo"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
