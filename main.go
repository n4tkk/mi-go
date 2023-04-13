package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/n4tkk/mi-go/core"
	"github.com/n4tkk/mi-go/service/admin"
	"github.com/n4tkk/mi-go/service/miauth"
)

var (
	miHost = "http://localhost:3000"
)

func main() {
	sessionId := uuid.New()
	client := core.NewClient(miHost+"/api/", "")

	authUrl := fmt.Sprint(miHost, "/miauth/", sessionId)
	fmt.Println(authUrl)

	fmt.Scanln()

	miauthSvc := miauth.NewService(client)
	token, err := miauthSvc.GetToken(sessionId.String())
	if err != nil {
		fmt.Println(err)
	}

	client.Token = token

	adminSvc := admin.NewService(client)
	serverInfo, err := adminSvc.GetServerInfo()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(serverInfo)
	}
}
