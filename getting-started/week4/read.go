package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	firstName string
	lastName string
}

func main() {
	printFileContent()
}

func printFileContent() {
	fileStream := getFileStream()
	defer fileStream.Close()
	scanner := bufio.NewScanner(fileStream)
	nameSlice := make([]name, 0, 20)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		nameSlice = append(nameSlice, name{words[0], words[1]})
	}
	for _, name := range nameSlice {
		fmt.Println(name.firstName, name.lastName)
	}
}

func getFileStream() *os.File {
	var fileName string
	fmt.Println("please enter the abs path of the .txt file: ")
	fmt.Scan(&fileName)
	fileStream, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return fileStream
}
