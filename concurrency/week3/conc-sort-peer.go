package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func softsilceby(wg *sync.WaitGroup, slice []int) {
	sort.Ints(slice)
	wg.Done()
}

func main() {
	var input []string
	fmt.Printf("Enter numbers separated by space then enter ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		input = strings.Split(line, " ")
	}

	numbers := make([]int, 0, len(input))
	for _, raw := range input {
		v, err := strconv.Atoi(raw)
		if err != nil {
			log.Print(err)
			continue
		}
		numbers = append(numbers, v)
	}

	var wg sync.WaitGroup
	length := len(numbers)
	incre := length / 4

	wg.Add(4)
	for j := 1; j <= 4; j++ {
		index := incre * j
		if j == 4 && (length%4) > 0 {
			index += (length % 4)
		}
		fmt.Println("Slice", j, numbers[(incre*(j-1)):index])
		go softsilceby(&wg, numbers[(incre*(j-1)):index])
	}
	wg.Wait()
	sort.Ints(numbers)
	fmt.Println("Sorted array: ", numbers)
}
