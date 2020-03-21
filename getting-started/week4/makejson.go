package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func getParam(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(message)
	param, err := reader.ReadString('\n')
	if (err== nil) {
		param = strings.TrimRight(param, "\r\n")
	}
	return param
}

func constructJson() {

	name := getParam("please enter a name:")
	address := getParam("please enter an address:")

	jsonObj := map[string]string{}
	jsonObj["name"] = name
	jsonObj["address"] = address

	jsonData, err := json.Marshal(jsonObj)
	if (err == nil) {
		fmt.Println(string(jsonData))
	} else {
		fmt.Println("error while marshalling to json")
	}
}

func main() {
	constructJson()
}


