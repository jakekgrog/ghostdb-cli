package commands

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command {
	Use:   "GhostDB",
	Short: "GhostDB Command-line Interface",
	Long:  "Interact and administrate your GhostDB cluster",
}

// Execute is called as the run operation for the RootCmd
// parameters: nil
// returns: nil
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}