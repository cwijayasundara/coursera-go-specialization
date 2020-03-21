package main

import "fmt"

func forLoop1()  {
	for i:=0; i<10; i++ {
		fmt.Println("Hi! there..from for loop")
	}
}

func forLoop2()  {
	i := 1
	for i < 10 {
		fmt.Println("Hi! there.. from do while style for loop")
		i++
	}
}

func forLoop3(){
	for {
		fmt.Println("Hi..")
	}
}

func main() {
	forLoop1()
	forLoop2()
	forLoop3()
}
