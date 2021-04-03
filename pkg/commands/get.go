package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/jakekgrog/ghostdb-cli/pkg/structures"
)

var GetCmd = &cobra.Command {
	Use: "get",
	Short: "Get a key/value pair from the store",
	Long: "Get a key/value pair from the store by asking any node in the cluster",
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		addr, _ := cmd.Flags().GetString("addr")

		if key == "" || addr == "" {
			fmt.Println("You must supply a value for:")
			fmt.Println(" - The key for the item you want to retrieve (e.g --key \"myKey\")")
			fmt.Println(" - The address of a node in the cluster (e.g --addr 127.0.0.1:7991)")
		} else {
			getPair(key, addr)
		}
	},
}

func init() {
	RootCmd.AddCommand(GetCmd)
	GetCmd.Flags().StringP("key", "k", "", "specify the key for the item you want to retrieve")
	GetCmd.Flags().StringP("addr", "a", "", "specify the address of a node in the cluster")
}

func getPair(key, addr string) {
	data := structures.NewStoreRequest(key, "")
	response := makePostRequest(addr+"/get", data)
	fmt.Println(response.Gobj.Value)
}