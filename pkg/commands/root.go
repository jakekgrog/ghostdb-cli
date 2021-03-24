package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/jakekgrog/ghostdb-cli/pkg/structures"
)

var (
	defaultPort = "7991"
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

func makePostRequest(url string, data structures.CacheRequest) structures.CacheResponse {
	client := &http.Client{}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	req, _ := http.NewRequest("POST", "http://"+url, reqBodyBytes)
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var responseObject = new(structures.CacheResponse)
	json.Unmarshal(bodyBytes, responseObject)

	return *responseObject
}
