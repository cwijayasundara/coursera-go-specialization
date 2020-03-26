package main

import (
	"fmt"
	"sort"
)

func acceptInput(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return acceptInput(x, err)
}

func sortSlice(slice []int, c chan []int) {
	fmt.Println(slice)
	sort.Sort(sort.IntSlice(slice))
	c <- slice
}

func splitToChunks(toSort []int) ([]int, []int, []int, []int) {
	chunk := len(toSort) / 4
	s1 := toSort[0:chunk]
	s2 := toSort[chunk : chunk*2]
	s3 := toSort[chunk*2 : chunk*3]
	s4 := toSort[chunk*3 : len(toSort)]
	return s1, s2, s3, s4
}

func main() {

	c := make(chan []int, 4)
	sorted := []int{}

	fmt.Println("pls enter a space delimited sequence of ints, then press space enter to execute::")
	fmt.Print("> ")
	arrToSort := acceptInput([]int{}, nil)

	// split to 4 chunks and execute all in parellel
	s1, s2, s3, s4 := splitToChunks(arrToSort)

	// add the processed slices to a global array to sort
	go sortSlice(s1, c)
	sorted1 := <-c
	go sortSlice(s2, c)
	sorted2 := <-c
	go sortSlice(s3, c)
	sorted3 := <-c
	go sortSlice(s4, c)
	sorted4 := <-c

	sorted = append(sorted, sorted1...)
	sorted = append(sorted, sorted2...)
	sorted = append(sorted, sorted3...)
	sorted = append(sorted, sorted4...)

	sort.Sort(sort.IntSlice(sorted))
	fmt.Println(sorted)
}
