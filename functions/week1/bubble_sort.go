package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	slice := readTenInts()
	bubbleSort(slice)
	fmt.Println(slice)
}

func bubbleSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		for j := 0; j < len(slice)-i; j++ {
			if slice[j] > slice[j+1] {
				swapValues(slice, j)
			}
		}
	}
}

func swapValues(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func readTenInts() []int {
	//define the number of array to 10
	size := 10
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter 10 ints to sort, pls enter after each int")
	intArray := make([]int, size)
	for count := 0; count < size; count++ {
		scanner.Scan()
		item, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("incorrect value entered !")
		}
		intArray[count] = item
	}
	return intArray
}
