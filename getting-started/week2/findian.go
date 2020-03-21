package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func filter(filterstr string) bool {
	isFound := false
	str := strings.ToLower(filterstr)
	if strings.HasPrefix(str, "i") && strings.Contains(str, "a") && strings.HasSuffix(str, "n") {
		isFound = true
	}
	return isFound
}

func main() {
	var name string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter the name :")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")
	name = strings.ReplaceAll(name, " ", "")
	charsFound := filter(name)
	if(charsFound) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

