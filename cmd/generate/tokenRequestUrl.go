/*
Copyright © 2023 Natsuki Kirigakure
*/
package generate

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/n4tkk/mi-go/service/miauth"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tokenRequestUrlCmd represents the tokenRequestUrl command
var tokenRequestUrlCmd = &cobra.Command{
	Use: "tokenRequestUrl",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("tokenRequestUrl called")
		sessionId := uuid.New().String()
		app := miauth.App{
			MisskeyHost: viper.GetString("host"),
		}
		fmt.Printf("host: %s\n", app.MisskeyHost)
		fmt.Printf("sessionId: %s\n", sessionId)
		fmt.Println("--------")
		fmt.Printf("url: %s\n", app.GenerateRequestUrl(sessionId))
	},
}

func init() {
	tokenRequestUrlCmd.PersistentFlags().StringP("host", "H", "https://misskey.io", "Misskey host")
	viper.BindPFlag("host", tokenRequestUrlCmd.PersistentFlags().Lookup("host"))
}
