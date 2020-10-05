package router

import (
	"encoding/json"
	"net/http"
	"time"
)

type UserResponse struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
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

func Error(w http.ResponseWriter, code int, message string) {
	body, _ := json.Marshal(map[string]string{"error": message})
	w.WriteHeader(code)
	w.Write(body)
}
