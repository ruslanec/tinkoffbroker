package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type rpcError struct {
	Message     string `json:"message,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

func main() {
	jsonFile, err := os.Open("./service/api_errors.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("loaded file: ", jsonFile.Name())
	defer jsonFile.Close()

	var rpc_errors map[string]rpcError
	content, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(content, &rpc_errors)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rpc_errors["40002"].Description)
}
