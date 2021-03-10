package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"runtime"

	"github.com/shirou/gopsutil/process"
	"github.com/spf13/cobra"
)

var NodesCmd = &cobra.Command {
	Use:   "nodes [add|get]",
	Short: "Interacts with ghostdb nodes",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		switch args[0] {
			case "get":
				getNodes()
			case "add":
				name, _ := cmd.Flags().GetString("name")
				http, _ := cmd.Flags().GetString("addr")
				raft, _ := cmd.Flags().GetString("raft")
				join, _ := cmd.Flags().GetString("join")
				
				if name == "" || http == "" || raft == "" {
					fmt.Println("You must supply values for:\n - The name of the new node (e.g -n worker1)\n - The address of the new node (e.g -a :7991)\n - The raft address of the new node (e.g -r :11000)")
				} else {
					addNodes(name, http, raft, join)
				}
		}
	},
}

func init() {
	RootCmd.AddCommand(NodesCmd)
	NodesCmd.Flags().StringP("name", "n", "", "specify the name of the new node (e.g worker1)")
	NodesCmd.Flags().StringP("addr", "a", "", "specify the address of the new node (e.g :7991)")
	NodesCmd.Flags().StringP("raft", "r", "", "specify the raft address of the new node (e.g :11000)")
	NodesCmd.Flags().StringP("join", "j", "", "specify the address of the leader node (only required when adding a follower node) (e.g :7991)")
}

type Data struct {
	error error
}

func getNodes() {
	pids, _ := process.Pids()
	fmt.Println(format("PID", "NAME", "HTTP PORT", "RAFT PORT"))
	var http, raft string
	for _, pid := range pids {
		proc, _ := process.NewProcess(pid)
		name, _ := proc.Name()
		var procName string
		if runtime.GOOS == "windows" {
			procName = "ghostdb.exe"
		} else {
			procName = "ghostdb"
		}
		if name == procName {
			cmdline, _ := proc.Cmdline()
			line := strings.Split(string(cmdline), " ")
			if len(line) > 6 {
				name, http, raft = line[2], line[4], line[6]
			} else {
				name = line[2]
				http = ":7991"
				raft = ":11000"
			}
			fmt.Println(format(fmt.Sprint(pid), name, http, raft))
		}
	}
}

func addNodes(name, http, raft, join string) {
	c := make(chan Data)
	if len(http[1:]) > 55 {
		fmt.Println("The port " + http + " does not exist.")
		os.Exit(2)
	}
	if len(raft[1:]) > 55 {
		fmt.Println("The port " + raft + " does not exist.")
		os.Exit(2)
	}
	go createNode(c, name, http, raft, join)
	res := <-c
	if res.error != nil {
		fmt.Println("Failed to execute command, only root may add more nodes")
	} else {
		fmt.Println("Node deployed!")
	}
}

func createNode(ch chan<- Data, name, http, raft, join string) {
	var cmd *exec.Cmd
	if join == "" {
		cmd = exec.Command("C:\\Program Files (x86)\\GhostDB\\ghostdb.exe", "-id", name, "-http", http, "-raft", raft, "./"+name, "&")
	} else {
		cmd = exec.Command("C:\\Program Files (x86)\\GhostDB\\ghostdb.exe", "-id", name, "-http", http, "-raft", raft, "-join", join, "./"+name, "&")
	}
	err := cmd.Start()
	ch <- Data {
		error: err,
	}
	
}

func format(pid string, name string, http string, raft string) string {
	return fmt.Sprintf("%-20s%-20s%-20s%-20s\n", pid, name, http, raft)
}