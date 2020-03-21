package main

import (
	"fmt"
	"sort"
	"strconv"
)

func process() {
	// initial slice
	intArray := make([]int, 3)
	// input as string
	var input string
	count := 0
	for {
		fmt.Printf("Enter an integer value: ")
		fmt.Scan(&input)

		if input == "X" {
			break
		}
        // convert the string to int
		if value, err := strconv.Atoi(input); err == nil {
			if count < 3 {
				intArray[count] = value
			} else {
				intArray = append(intArray, value)
			}
			sortedIntArray := make([]int, len(intArray))
			copy(sortedIntArray, intArray)
			sort.Ints(sortedIntArray)
			fmt.Println(sortedIntArray)
			count++
		}
	}
}

func main() {
	process()
}


