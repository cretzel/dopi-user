package router

import (
	"dopi-user/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	MuxRouter   *mux.Router
	userService model.UserService
}

func NewUserRouter(userService model.UserService) *UserRouter {

	muxRouter := mux.NewRouter()
	amw := AuthMiddleware{}
	muxRouter.Use(amw.authMiddleware)
	userRouter := UserRouter{MuxRouter: muxRouter, userService: userService}

	muxRouter.HandleFunc("/api/user/login", userRouter.PostLogin).Methods("POST")
	muxRouter.HandleFunc("/api/user/refresh", userRouter.PostRefresh).Methods("POST")
	muxRouter.HandleFunc("/api/user/logout", userRouter.PostLogout).Methods("POST")
	muxRouter.HandleFunc("/api/user/users/me", userRouter.GetMe).Methods("GET")
	muxRouter.HandleFunc("/api/user/users/{username}", userRouter.GetUser).Methods("GET")
	muxRouter.HandleFunc("/api/user/users/{username}", userRouter.UpdateUser).Methods("PUT")
	muxRouter.HandleFunc("/api/user/users", userRouter.GetUsers).Methods("GET")

	return &userRouter
}

func (ur *UserRouter) PostLogin(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		Error(w, 400, "Invalid body")
		return
	}

	user, err := ur.userService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		Error(w, 400, "Invalid credentials")
		return
	}

	ur.createJwtReponse(user, w)
}

func (ur *UserRouter) PostRefresh(w http.ResponseWriter, r *http.Request) {
	claims := ur.claims(r)
	user, err := ur.userService.GetUserByUsername(claims.Username)
	if err != nil {
		Error(w, 404, "Not Found")
		return
	}

	ur.createJwtReponse(user, w)
}

func (ur *UserRouter) PostLogout(w http.ResponseWriter, r *http.Request) {
	jwtCookie := http.Cookie{
		Name:     "dopi_jwt",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(w, &jwtCookie)
}

func (ur *UserRouter) createJwtReponse(user *model.User, w http.ResponseWriter) {
	tokenString, _, claims, err := CreateJwt(user)
	if err != nil {
		Error(w, 500, "Internal server error (jwt)")
		return
	}
	jwtCookie := http.Cookie{
		Name:     "dopi_jwt",
		Value:    tokenString,
		HttpOnly: true,
	}
	http.SetCookie(w, &jwtCookie)

	response := UserInfoResponse{
		Username:  user.Username,
		Roles:     claims.Roles,
		ExpiresAt: time.Unix(claims.ExpiresAt, 0),
		IssuedAt:  time.Unix(claims.IssuedAt, 0),
	}
	json.NewEncoder(w).Encode(response)
}

func (ur *UserRouter) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := ur.userService.GetUserByUsername(params["username"])
	if err != nil {
		Error(w, 404, "Not Found")
		return
	}

	response := toUserDto(user)
	json.NewEncoder(w).Encode(response)
}

func (ur *UserRouter) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := ur.userService.GetUsers()
	if err != nil {
		Error(w, 404, "Not Found")
		return
	}

	response := toUserDtos(users)
	json.NewEncoder(w).Encode(response)
}

func (ur *UserRouter) GetMe(w http.ResponseWriter, r *http.Request) {
	claims := ur.claims(r)
	user, err := ur.userService.GetUserByUsername(claims.Username)
	if err != nil {
		Error(w, 404, "Not Found")
		return
	}

	response := toUserDto(user)
	json.NewEncoder(w).Encode(response)
}

func (ur *UserRouter) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userDto UserDto
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		Error(w, 400, "Bad Request")
		return
	}

	user, err := ur.userService.GetUserByUsername(userDto.Username)
	if err != nil {
		Error(w, http.StatusNotFound, fmt.Sprintf("User not found: %s", userDto.Username))
		return
	}

	user.Roles = userDto.Roles
	newUser, err := ur.userService.UpdateUser(user)
	json.NewEncoder(w).Encode(toUserDto(newUser))
}

func (ur *UserRouter) claims(r *http.Request) DopiClaims {
	return (r.Context().Value("claims")).(DopiClaims)
}

func (ur *UserRouter) dumpRequest(r *http.Request) {
	requestDump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(requestDump))
}
