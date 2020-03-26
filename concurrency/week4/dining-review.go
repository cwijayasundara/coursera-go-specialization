package main

import (
	"fmt"
	"sync"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	leftCS, rightCs *ChopStick
	idx             int
}

type Host struct {
	count int
}

var host Host

func (h *Host) request(wgp *sync.WaitGroup) {
	for h.count > 1 {
	}
	h.count++
	wgp.Done()
}

func (h *Host) free() {
	h.count--
}

func (p *Philosopher) eat() {
	for i := 0; i < 3; i++ {
		var wgp sync.WaitGroup
		wgp.Add(1)
		go host.request(&wgp)
		wgp.Wait()

		p.leftCS.Lock()
		p.rightCs.Lock()

		fmt.Println("starting to eat ", p.idx)
		fmt.Println("finshed eating ", p.idx)

		p.rightCs.Unlock()
		p.leftCS.Unlock()

		host.free()
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{chopSticks[i], chopSticks[(i+1)%5], i + 1}
	}

	host = Host{0}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat()
	}
	wg.Wait()
}
