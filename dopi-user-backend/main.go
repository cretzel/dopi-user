package main

import (
	"dopi-user/router"
	"dopi-user/service"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load()

	session, err := service.NewSession()
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}
	defer session.Close()

	userService := service.NewUserService(session)
	userRouter := router.NewUserRouter(userService)

	port := os.Getenv("PORT")
	log.Println("Starting server on the port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, userRouter.MuxRouter))
}
