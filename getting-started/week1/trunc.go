package main

import (
	"fmt"
	"math"
	"os"
)

func acceptFloat () {
	var floatNum float64
	fmt.Println("Please enter a floting point number")
	 _, err := fmt.Fscan(os.Stdin, &floatNum)
	 if(err == nil) {
		 fmt.Println("Float num is " ,floatNum)
		 fmt.Println("After droping the decimal values", math.Floor(floatNum))
	 } else {
		 fmt.Println(err)
	 }
}

func main() {
	acceptFloat()
}
