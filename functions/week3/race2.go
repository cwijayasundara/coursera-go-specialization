/*
 * Race condition: A race condition is when two or more routines have access to the same/shared resource, such as a variable or
 * data structure and attempt to read and write to that resource without any regard to the other routines.
 * In the below example "count" is a mutable global veriable.
 * when I run this few times I see some times the output is x:1 y:0, x:0 y:1 etc. The output is not determinstic.
 */
package main

import (
	"fmt"
	"time"
)

var x, y int

func printY() {
	x = 1
	fmt.Print("y:", y, " ")
}

func printX() {
	y = 1
	fmt.Print("x:", x, " ")
}

func main() {
	go printY()
	go printX()
	time.Sleep(1 * time.Second)
}
