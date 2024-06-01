package main

import (
	"fmt"
	"sync"
)

const (
	numPhilosophers = 5
	numEats         = 3
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number         int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
	host           *Host
}

type Host struct {
	permission chan struct{}
}

func (h *Host) allowEating() {
	h.permission <- struct{}{}
}

func (h *Host) stopEating() {
	<-h.permission
}

func (p *Philosopher) eat() {
	for i := 0; i < numEats; i++ {
		p.host.allowEating()

		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Printf("Starting to eat %d\n", p.number)
		fmt.Printf("Finishing eating %d\n", p.number)

		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()

		p.host.stopEating()
	}
}

func main() {
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	host := &Host{permission: make(chan struct{}, 2)}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			number:         i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numPhilosophers],
			host:           host,
		}
	}

	var wg sync.WaitGroup
	for _, p := range philosophers {
		wg.Add(1)
		go func(p *Philosopher) {
			defer wg.Done()
			p.eat()
		}(p)
	}
	wg.Wait()
}
