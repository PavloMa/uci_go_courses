package main

/*Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each 
adjacent pair of philosophers.

Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not lowest-numbered first 
(which we did in lecture).

In order to eat, a philosopher must get permission from a host which executes in its own 
goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints 
“starting to eat <number>” on a line by itself, where <number> is the number of the 
philosopher.

When a philosopher finishes eating (before it has released its locks) it prints 
“finishing eating <number>” on a line by itself, where <number> is the number of 
the philosopher.
*/

/* Solution:
1. Before picking up the chopsticks, a philosopher must ask permission from the host. 
The host gives permission to only one philosopher at a time 
so that the philosopher has picked up both of their chopsticks. 

2. If one philosopher is eating, two adjacent philosophers cannot eat. Not enough chopsticks.

3. There are 5 philosophers and 5 chopsticks. 4 chopsticks is enough for 2 philosophers
to eat simultaneously. Integral diviosion of 5 through 2.

4. Host-philosopher communications are effected through channels.
A single channel is shared by all philosophers to request eating permission
Five separate channels (one per philosopher) used for sending permission from the host

6. Philosopher -> Host communication is effected through one channel:
The host is listening to the single channel for:
- incoming philosophers' Ids (positive int) asking for permission to eat. 
- If -1 is received, it means some philosopher has stopped eating. 

7. Host -> Philosopher communication is effected through 5 channels
The Philosopher is listening to its individual channel (1 of 5)
- If a Philosopher has eaten 3 times, the host sends a -1, which makes the 
philosopher terminate the request-eat loop (stop asking for eat permission). 
- If less than 2 Philosophers are eating, then the host allows the philosopher 
to eat by sending him back +1. 
- If 2 Philosophers are already eating then the host sends back a 0, 
so that the philosopher would not take any action.

8. Host algorithms:
- If an Id is received by the host then the host checks if the philosopher 
has already eaten 3 times. 
- If he has, the host sends a -1, which makes the philosopher terminate the request-eat loop 
(stop asking for eating permission). 
- If the philosopher has not eaten 3 times, then the host additionally checks 
how many philosophers are currently eating. 
- If less than 2 are eating, then the host allows the philosopher to eat by sending back +1. 
- If 2 are already eating then the host sends back a 0, so that the philosopher would not take an action.

9. P.S. It is not obvious, who controls the number of eats: the host of a philosopher.
So, we include the controls on the part of the host (it controls the whole process)
*/

import (
	"sync"
	"time"
    "math/rand"
    "fmt"
)

type Chopstick struct { 
	sync.Mutex // Go will automatically use the type as the name of the anonymous field
}

type Philosopher struct {
	number int // philosopher id
	eaten int // number of time the philosopher has eaten
	leftChopstick, rightChopstick *Chopstick
}

type Host struct {
	eating int // number of philosophers eating now
	finished int // number of philiosophers, who finished eating
	eatingTable map[int]int //who has eaten how many times (id -> number of eats)
}

// primary task parameters
const quantity int = 5 // number of philosophers and chopsticks
const parallelism int = quantity / 2 // 5 through 2 is 2 philosophers allowed to eat simultaneousl	

const eatingRounds int = 3 // number of simulated eatings for each philosopher

// mesage codes:
// host -> philosopher
const eatPermission int = 1
const stopEating int = -1
const idle = 0
// philosopher -> host
const finishedEating int = -1


// The function simulates eating by a philosopher
func (p *Philosopher) eat() {
	
	if (*p).eaten < eatingRounds { // hungry yet
		if RandBool() { // randomly pick up sticks
			(*p).leftChopstick.Lock()
			(*p).rightChopstick.Lock()
		} else {
			(*p).rightChopstick.Lock()
			(*p).leftChopstick.Lock()
		}

	    fmt.Println("Starting to eat ", (*p).number)
		time.Sleep(1 * time.Second) // eating tekes up some time
	    fmt.Println("Finished eating ", (*p).number)
	    (*p).eaten++

		if RandBool() { // randomly lay down sticks
			(*p).leftChopstick.Unlock()
			(*p).rightChopstick.Unlock()
		} else {
			(*p).rightChopstick.Unlock()
			(*p).leftChopstick.Unlock()
		}
	}

}

func (p *Philosopher) askToEat(wg *sync.WaitGroup, toHost chan int, fromHost chan int) {

	// for ; (*p).eaten < eatingRounds; { // try to get permission to eat until explicitly rejected by the host
	for { // try to get permission to eat until explicitly rejected by the host
		 
		toHost <- (*p).number // Notify the host the philosopher wants to eat
		
		message := <- fromHost // Wait for a response from the host

		switch message {
			case eatPermission:
				(*p).eat() // If allowed, eat 
				toHost <- finishedEating // When the philosopher finishes one eat, he sends -1 to the host
			case stopEating:
				wg.Done()
				return
		}
	}
	// wg.Done() // stops eating itself, while he knows he has reached 3
}

func (h *Host) manageEating(wg *sync.WaitGroup, toHost chan int, fromHost []chan int) {
	for {
		// every time someone has eaten 3 times, we increase by 1,
		// so, when we reach 5, the host has done its job
		if (*h).finished >= quantity {
			wg.Done()
			return
		}
		// read which philosopher requests permission, listen at channel 5
		// for incoming requests
		result := <- toHost // listen to incoming philosopher permission requests 
							// or info that a philosopner finished eating  

		if result == finishedEating { // a philosopher sent -1, it means he has stopped eating
			(*h).eating -= 1 
		} else { // result is a philosopher's id, requesting an eat

			id := result
			if (*h).eatingTable[id] >= eatingRounds { 
				(*h).finished += 1 // one philosopher more has finished eating
				fromHost[id-1] <- stopEating // if the philosopher has eaten 3 times, send him -1 to stop
			} else { // this philosopher is still allowed to eat
				if (*h).eating < parallelism { // if there are less than 2 philosophers eating
					(*h).eatingTable[id] += 1 // was allowed to eat one extra time
					fromHost[id-1] <- eatPermission
					(*h).eating += 1 // increment the counter
				} else { // 2 philosphers are already eating, send 0 so that the philosopher idles
					fromHost[id-1] <- idle
				}
			}

		}
	}
}

// The function returns a random boolean value based on the current time
func RandBool() bool {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(2) == 1
}


func CheckFull(philosophers []Philosopher) bool {
	
	var hungry bool = false // assume no one is hungry, i.e has eaten 3 times
	for i := 1; i <= quantity; i++ {
	   if philosophers[i-1].eaten < eatingRounds { // if any has actually eaten not yet 3 times
	   		hungry = true
	   		return hungry
	   	}
	}
	return hungry
}

func main() {

	var host Host
	var chopsticks []Chopstick = make([]Chopstick, quantity)
	var philosophers []Philosopher = make([]Philosopher, quantity)

	var wg sync.WaitGroup

	var toHost chan int = make(chan int) // one channel Philosopher -> Host
	var fromHost[] chan int = make([]chan int, quantity) // 5 channels Host -> Philosophers 
	for i :=0 ; i < quantity; i++ { 
	 	fromHost[i] = make(chan int)
	}

	for i := 0; i < quantity; i++ {
		chopsticks[i] = Chopstick{}
	}

	for i := 0; i < quantity; i++ {
	   philosophers[i] = Philosopher{i + 1, 0, &chopsticks[i], &chopsticks[(i+1)%5]}
	}

	host.finished = 0
	host.eating = 0
	host.eatingTable = map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
	
	wg.Add(quantity + 1) // philosophers + host

	go host.manageEating(&wg, toHost, fromHost) // run the host
	
	for i := 0; i < quantity; i++ { // run the philosophers
		go (philosophers[i]).askToEat(&wg, toHost, fromHost[i])
	}

	wg.Wait()
	return

}