/*
Copyright © 2023 Natsuki Kirigakure
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd represents the get command
var Cmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
