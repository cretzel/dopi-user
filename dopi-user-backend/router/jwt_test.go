package router

import (
	"dopi-user/model"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func Test_CreateJwtToken(t *testing.T) {
	t.Run("CreateJwtToken", createJwtToken_should_create_token)
}

func createJwtToken_should_create_token(t *testing.T) {
	//Arrange
	os.Setenv("JWT_TOKEN_VALIDITY", fmt.Sprint("", 365*24*60))

	username := "admin"
	roles := []string{"ROLE_ADMIN", "ROLE_USER", "ROLE_APP"}
	user := &model.User{Username: username, Roles: roles}
	secret := "asdf"
	t.Log("Secret", secret)

	//Act
	tokenString, _, claims, _ := CreateJwt(user)
	t.Log("Token", tokenString)
	t.Log("Claims", claims)
	t.Log("ExpiresAt", time.Unix(claims.ExpiresAt, 0))

	//Assert
	if tokenString == "" {
		log.Fatalf("Unable to create token")
		t.Error("Unable to create token")
	}
}
