/*
 * Race condition: A race condition is when two or more routines have access to the same/shared resource, such as a variable or
 * data structure and attempt to read and write to that resource without any regard to the other routines.
 * In the below example "count" is a mutable global veriable.
 * The below program has few go routies that mutate the global / shared veriables.
 * when I run this few times I can see the output is not determinstic.
 */

package main

import (
	"fmt"
	"time"
)

var count int
var x, y int

func race() {
	count++
	fmt.Println(count)
}

func printY() {
	x = 1
	fmt.Print("y:", y, " ")
}

func printX() {
	y = 1
	fmt.Print("x:", x, " ")
}

func main() {
	go race()
	go race()
	go printY()
	go printX()
	time.Sleep(1 * time.Second)
}
