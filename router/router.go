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
	muxRouter.Use(commonMiddleware)
	userRouter := UserRouter{MuxRouter: muxRouter, userService: userService}

	muxRouter.HandleFunc("/api/user/login", userRouter.PostLogin).Methods("POST")
	muxRouter.HandleFunc("/api/user/refresh", userRouter.PostRefresh).Methods("POST")
	muxRouter.HandleFunc("/api/user/logout", userRouter.PostLogout).Methods("POST")
	muxRouter.HandleFunc("/api/user/users/me", userRouter.GetMe).Methods("GET")
	muxRouter.HandleFunc("/api/user/users/{username}", userRouter.GetUser).Methods("GET")

	return &userRouter
}

func (ur *UserRouter) PostLogin(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginRequest)
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
	claims, err := ur.checkAuth(r)
	if err != nil {
		Error(w, 401, "Not authenticated")
		return
	}

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
	_, err := ur.checkAuth(r)
	if err != nil {
		Error(w, 401, "Not authenticated")
		return
	}

	params := mux.Vars(r)

	user, err := ur.userService.GetUserByUsername(params["username"])
	if err != nil {
		Error(w, 404, "Not Found")
		return
	}

	response := toUserResponse(user)
	json.NewEncoder(w).Encode(response)
}

func (ur *UserRouter) GetMe(w http.ResponseWriter, r *http.Request) {
	dopiClaims, err := ur.checkAuth(r)
	if err != nil {
		Error(w, 401, "Not authenticated")
		return
	}

	user, err := ur.userService.GetUserByUsername(dopiClaims.Username)
	if err != nil {
		Error(w, 404, "Not Found")
		return
	}

	response := toUserResponse(user)
	json.NewEncoder(w).Encode(response)
}

func (ur *UserRouter) checkAuth(r *http.Request) (*DopiClaims, error) {
	jwtCookie, err := r.Cookie("dopi_jwt")
	if err != nil {
		return nil, err
	}

	token, err := VerifyToken(jwtCookie.Value)
	if err != nil {
		return nil, err
	}

	dopiClaims, err := GetClaims(token)
	if err != nil {
		return nil, err
	}

	return dopiClaims, nil
}

func (ur *UserRouter) dumpRequest(r *http.Request) {
	requestDump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(requestDump))
}
