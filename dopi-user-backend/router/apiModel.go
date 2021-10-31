package router

import (
	"dopi-user/model"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type UserDto struct {
	Id       string   `json:"id,omitempty"`
	Username string   `json:"username,omitempty"`
	Roles    []string `json:"roles"`
}

type CreateUserDto struct {
	Username string   `json:"username,omitempty"`
	Roles    []string `json:"roles"`
	Password string   `json:"password,omitempty"`
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

func toUserDto(user *model.User) *UserDto {
	return &UserDto{
		Username: user.Username,
		Roles:    user.Roles,
	}
}

func toUser(user *UserDto) *model.User {
	return &model.User{
		Username: user.Username,
		Roles:    user.Roles,
	}
}

func createUserDtoToUser(user *CreateUserDto) *model.User {
	return &model.User{
		Username: user.Username,
		Roles:    user.Roles,
		Password: user.Password,
	}
}

func toUserDtos(users []model.User) []UserDto {
	userArray := make([]UserDto, len(users))
	for i, u := range users {
		userArray[i] = *toUserDto(&u)
	}
	return userArray
}

func Error(w http.ResponseWriter, code int, message string) {
	body, _ := json.Marshal(map[string]string{"error": message})
	w.WriteHeader(code)
	w.Write(body)
}

func StatusError(w http.ResponseWriter, err error) {
	log.Println(err)
	reason := http.StatusText(http.StatusInternalServerError)
	code := http.StatusInternalServerError

	switch err.(type) {
	case *model.StatusError:
		reason = err.(*model.StatusError).Reason
		code = err.(*model.StatusError).Code
	}
	body, _ := json.Marshal(map[string]string{"error": reason})
	w.WriteHeader(code)
	w.Write(body)
}
