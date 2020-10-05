package router

import (
	"dopi-user/model"
	"encoding/json"
	"net/http"
	"time"
)

type UserResponse struct {
	Id       string   `json:"id,omitempty"`
	Username string   `json:"username,omitempty"`
	Roles    []string `json:"roles"`
}

type UserInfoResponse struct {
	Username  string    `json:"username"`
	Roles     []string  `json:"roles"`
	ExpiresAt time.Time `json:"expiresAt"`
	IssuedAt  time.Time `json:"issuedAt"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func toUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		Username: user.Username,
		Roles:    user.Roles,
	}
}

func Error(w http.ResponseWriter, code int, message string) {
	body, _ := json.Marshal(map[string]string{"error": message})
	w.WriteHeader(code)
	w.Write(body)
}
