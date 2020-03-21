package main

import "fmt"

func caseTest(count int32) {
	switch count {
	case 1 :
		fmt.Println("case 1")
	case 2 :
		fmt.Println("case 2")
	default:
		fmt.Println("default case")
	}
}

func taglessCase (x int32) {
	switch {
	case x > 0 :
		fmt.Println("case positive")
	case x < 0 :
		fmt.Println("case nagative")
	default:
		fmt.Println("no case")
	}
}

func main() {
	caseTest(1)
	taglessCase(2)
}

