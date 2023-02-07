package main

import (
	"fmt"
	"time"
)

/*
*count* will be our shared variable
 */
var count int

/*
The *increment* function increases the shared variable
without any measure of synchronization.
*/
func increment() {
	for i := 0; i < 100; i++ {
		fmt.Println("incrementing, before : ", count)
		count = count + 1
		fmt.Println("incrementing, after : ", count)
	}
}

/*
The *decrement* function decreases the shared variable
without any measure of synchronization.
*/
func decrement() {
	for i := 0; i < 100; i++ {
		fmt.Println("decrementing, before : ", count)
		count = count + 1
		fmt.Println("decrementing, after : ", count)
	}
}

func main() {
	/*
		When executing both functions as Go routines, it is
		not possible to know for how long each routine will be
		executed, nor what will be the value of the shared
		variable. This is due to the problem of interleaving in
		the execution of the functions.
	*/
	go increment()
	go decrement()

	time.Sleep(time.Second)
	fmt.Println("done")
}
