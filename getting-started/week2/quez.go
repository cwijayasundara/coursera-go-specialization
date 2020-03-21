package main

import "fmt"

func main() {
	var k int = 10
	k = 100
	println("Variable k =", k)

	for i:=1; i<4; i++ {

		for j := 1; j<6; j++ {
			fmt.Println("Value of i =", i)
			fmt.Println("Value of j =", j)
			fmt.Println("I + J = " , i+j)
			fmt.Println("+++++++")
		}
	}
	m := 7
	fmt.Println(m*m)
}