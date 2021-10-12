package main

import (
	"dopi-user/router"
	"dopi-user/service"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load()

	client, err := service.NewClient()
	if err != nil {
		log.Fatalln("unable to connect to mongodb", err)
	}
	defer client.Close()

	userService := service.NewUserService(client)
	userRouter := router.NewUserRouter(userService)

	port := os.Getenv("PORT")
	log.Println("Starting server on the port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, userRouter.MuxRouter))
}
