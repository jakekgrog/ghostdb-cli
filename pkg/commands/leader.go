package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/jakekgrog/ghostdb-cli/pkg/structures"
)

var LeaderCmd = &cobra.Command {
	Use: "leader",
	Short: "Find the leader of a cluster",
	Long: "Find the leader of a cluster by asking any node within in the cluster",
	Run: func(cmd *cobra.Command, args []string) {
		addr, _ := cmd.Flags().GetString("addr")
		if addr == "" {
			fmt.Println("You must supply a value for:\n - The address of a known node in the cluster (e.g. 127.0.0.1:7991)")
		} else {
			ldr := getLeader(addr)
			fmt.Println(ldr)
		}
	},
}

func init() {
	RootCmd.AddCommand(LeaderCmd)
	LeaderCmd.Flags().StringP("addr", "a", "", "specify the address of a known node in a cluster (e.g 127.0.0.1:7991)")
}

func getLeader(addr string) string {
	data := structures.NewEmptyRequest()
	response := makePostRequest(addr+"/getLeader", data)
	ldrAddrPort := response.Message
	ldrAddr := strings.Split(ldrAddrPort, ":")[0]
	return ldrAddr+":"+defaultPort
}
