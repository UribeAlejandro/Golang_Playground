package main

import (
	"fmt"
	"sync"
)

var p int                        //Number of philosopher on the table
var n int                        //Number of times each philosopher is allowed to eat
var philosEatingMutex sync.Mutex //To block number of concurrent philosophers eating
var philosEating int

func main() {
	var wg sync.WaitGroup
	p = 5
	n = 2

	c := make(chan bool) //Creates a channel to ask for permission to eat and send back allowance

	CSticks := make([]*ChopStick, p)
	for i := 0; i < p; i++ {
		CSticks[i] = new(ChopStick)
	}

	Philos := make([]*Philo, p)
	for i := 0; i < p; i++ {
		Philos[i] = &Philo{0, i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	for i := 0; i < p; i++ {
		var allowed bool
		for allowed == true {
			wg.Add(1)
			go host(c, &wg)
			allowed = <-c
		}
		wg.Add(1)
		go Philos[i].eat(&wg)
	}
	wg.Wait()
}

func (p Philo) eat(wg *sync.WaitGroup) {
	for p.timesEaten < n {

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("Starting to eat ", p.id)

		fmt.Println("Finished eating ", p.id)
		p.rightCS.Unlock()
		p.leftCS.Unlock()

		p.timesEaten++
		philosEatingMutex.Lock()

		if philosEating > 0 {
			philosEating--
		}

		philosEatingMutex.Unlock()
	}
	defer wg.Done()

}

//Checks only two Philosophers are eating at the same time, gives permission and send it back through same channel
func host(c chan bool, wg *sync.WaitGroup) {

	allowed := false

	philosEatingMutex.Lock()

	//Only allows two Philosophers to eat at the same time
	if philosEating <= n {
		philosEating++
		allowed = true
	}

	philosEatingMutex.Unlock()

	c <- allowed
	wg.Done()
}

//Two objects are needed: Chopsticks and Philosophers

type ChopStick struct {
	sync.Mutex
}

type Philo struct {
	timesEaten      int
	id              int
	leftCS, rightCS *ChopStick
}
