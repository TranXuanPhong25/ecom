package models

type LoginResponse struct {
	Token string   `json:"token"`
	User  UserInfo `json:"userInfo"`
}
type LoginRequest struct {
	Credentials
}
