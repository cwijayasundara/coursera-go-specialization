package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopSticks struct {
	sync.Mutex
}

type Philosopers struct {
	leftCs  *ChopSticks
	rightCs *ChopSticks
	number  int
}

func (p *Philosopers) Eat(wg *sync.WaitGroup, host chan int) {
	for i := 0; i < 3; i++ {
		host <- p.number
		p.leftCs.Lock()
		p.rightCs.Lock()

		fmt.Println("starting to eat <", p.number, ">")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("finished eating <", p.number, ">")

		p.leftCs.Unlock()
		p.rightCs.Unlock()
		<-host
	}
	wg.Done()
}

func assignChops(philos []*Philosopers, i int, chops []*ChopSticks) {
	philos[i] = new(Philosopers)
	philos[i].number = i + 1
	philos[i].leftCs = chops[i]
	philos[i].rightCs = chops[(i+1)%5]
}

func main() {
	// create 5 chop sticks
	sticks := make([]*ChopSticks, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(ChopSticks)
	}
	// create 5 philosopers
	philos := make([]*Philosopers, 5)
	for i := 0; i < 5; i++ {
		assignChops(philos, i, sticks)
	}

	var wg sync.WaitGroup
	wg.Add(5)

	host := make(chan int, 2)
	for i := 0; i < 5; i++ {
		go philos[i].Eat(&wg, host)
	}
	wg.Wait()
}
