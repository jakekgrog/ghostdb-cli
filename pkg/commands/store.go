package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/jakekgrog/ghostdb-cli/pkg/structures"
)

var StoreCmd = &cobra.Command {
	Use: "store [put|add]",
	Short: "Put a value in the store",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		key, _   := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")
		addr, _  := cmd.Flags().GetString("addr")
		ldr, _   := cmd.Flags().GetBool("leader")
		
		if key == "" || value == "" || addr == ""{
			fmt.Println("You must supply a value for:")
			fmt.Println(" - A string value for the key of item you are storing (e.g --key \"myKey\")")
			fmt.Println(" - The value you want to store (e.g --value \"myValue\")")
			fmt.Println(" - The address of a node in the cluster (e.g --addr 127.0.0.1:7991)")
			fmt.Println("   - If --leader flag is not supplied, it must be the clusters leader address")
		} else {
			storeValue(args[0], key, value, addr, ldr)
		}
	},
}

func init() {
	RootCmd.AddCommand(StoreCmd)
	StoreCmd.Flags().StringP("key", "k", "", "specify the key for the item")
	StoreCmd.Flags().StringP("value", "v", "", "specify the value for the item")
	StoreCmd.Flags().StringP("addr", "a", "", "specify the address of a node in the cluster (e.g 127.0.0.1:7991)")
	StoreCmd.Flags().BoolP("leader", "l", false, "Allows user to specify the addr of any node and a background call will be made to find the leader")
}

func storeValue(cmd, key, value, addr string, ldr bool) {
	if !ldr {
		// Write to the node at addr.
		data := structures.NewStoreRequest(key, value)
		response := makePostRequest(addr+"/"+cmd, data)
		fmt.Println(response.Message)
	} else {
		// get the leader first, then put the value to that node
		ldrAddr := getLeader(addr)
		data := structures.NewStoreRequest(key, value)
		response := makePostRequest(ldrAddr+"/"+cmd, data)
		fmt.Println(response.Message)
	}
}