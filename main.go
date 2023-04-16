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

	app := miauth.App{
		MisskeyHost: "http://localhost:3000",
		Name:        "mi-go",
	}
	u := app.GenerateRequestUrl(sessionId.String())
	fmt.Println(u)

	fmt.Scanln()

	miauthSvc := miauth.NewService(client)
	token, err := miauthSvc.GetToken(sessionId.String())
	if err != nil {
		fmt.Println(err)
	}

	client.Token = token
	fmt.Println(client.Token)

	adminSvc := admin.NewService(client)
	serverInfo, err := adminSvc.GetServerInfo()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(serverInfo)
	}
}
