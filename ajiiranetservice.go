package main

import (
	"ajiiranetservice/constants"
	"ajiiranetservice/graph"
	"ajiiranetservice/vo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var ajiiraNetGraph = &graph.AjiiraGraph{}
var message = make(map[string]map[string]string, 0)

func main() {

	ajiiraNetGraph.ConnectedDevices = make(map[string]*vo.AJIIRADevice, 0)
	ajiiraNetGraph.Connections = make(map[string][]string, 0)

	// loadTestData() // To load the test data
	loadMessageData() // To load the messages

	var ajiiraNetHandler AJIIRANetHandler

	http.HandleFunc("/ajiranet/process", ajiiraNetHandler.ProcessData)

	fmt.Println("service started...!", constants.AjiraNetServicePortNo)

	err := http.ListenAndServeTLS(":"+fmt.Sprintf("%s", constants.AjiraNetServicePortNo), "certs"+string(os.PathSeparator)+"server.crt", "certs"+string(os.PathSeparator)+"server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func loadMessageData() {
	byts, err := ioutil.ReadFile("config" + string(os.PathSeparator) + "message.json")
	if err == nil {
		err = json.Unmarshal(byts, &message)
		if err != nil {
			fmt.Println("Cannot load the test data :", err.Error())
		}
	}
}

func loadTestData() {
	byts, err := ioutil.ReadFile("test" + string(os.PathSeparator) + "test.json")
	if err == nil {
		err = json.Unmarshal(byts, ajiiraNetGraph)
		if err != nil {
			fmt.Println("Cannot load the test data :", err.Error())
		}
	}

}
