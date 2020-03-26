package main

import (
	"fmt"
	"sync"
)

var i int = 0

var mut sync.Mutex

func inc() {
	mut.Lock()
	i = i + 1
	mut.Unlock()
}

func main() {
	go inc()
	go inc()
	fmt.Println(i)
}
