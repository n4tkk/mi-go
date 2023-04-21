package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/n4tkk/mi-go/core"
	"github.com/n4tkk/mi-go/service/admin"
	"github.com/n4tkk/mi-go/service/miauth"
	"github.com/n4tkk/mi-go/usecase"
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

	token, err := usecase.GetToken(miHost, sessionId.String())

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
