package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var PingCmd = &cobra.Command {
	Use: "ping",
	Short: "Ping Pong",
	Run: func(cmd *cobra.Command, args[]string) {
		fmt.Println("PONG!")
	},
}

func init() {
	RootCmd.AddCommand(PingCmd)
}