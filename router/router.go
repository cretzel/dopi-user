package router

import (
	"dopi-user/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"

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

	muxRouter.HandleFunc("/api/user/{username}", userRouter.GetUser).Methods("GET")
	muxRouter.HandleFunc("/api/user/login", userRouter.PostLogin).Methods("POST")

	return &userRouter
}

func (ur *UserRouter) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user, err := ur.userService.GetUserByUsername(params["username"])
	if err != nil {
		Error(w, 404, "Not Found")
		return
	}

	response := UserResponse{Username: user.Username}
	json.NewEncoder(w).Encode(response)
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

	response := UserResponse{Username: user.Username}
	json.NewEncoder(w).Encode(response)
}

func (ur *UserRouter) dumpRequest(r *http.Request) {
	requestDump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(requestDump))
}
