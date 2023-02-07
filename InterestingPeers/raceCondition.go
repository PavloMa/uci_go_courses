package main

import "fmt"

func main() {
	fmt.Println("Racecondition demonstration - the counter variable is updated by the increment and decrement functions.")
	fmt.Println("The functions are launched as two goroutines. If the routines were run in sequence they would always produce the same output of 0.")
	fmt.Println("The goroutines are sharing the same variable, this creates the potential for a race condition.")
	fmt.Println("The goroutines are launched on seperate threads. This allows the modification of the counter variable to occur out of order and there is no means to sychronize the modufication of the variable, between threads.")
	fmt.Println("The program is creating the conditions for a data race without any way of knowing the order of the modification of the counter variable.")
	fmt.Println("In addition the loop that launches the goroutines, can exit before the goroutines have finished executing, preventing the counter from being incremented and decremented an equal number of times.")
	fmt.Println("Goroutines are launched and then the program continues onto the next instruction, they do no block untill the routine has completed, which necessary for concurrent excecution of routines.")
	fmt.Println("This is due to the lack of a synchronization mechanism between the goroutines combined with the fact that it is sharing a common variable.")
	var counter int = 0
	for i := 0; i < 1000; i++ {
		go incrementCounter(&counter)
		go decrementCounter(&counter)
	}
	fmt.Println("\nCounter:", counter)
	fmt.Println("\nThe result would be 0 without a race condition. Run again and get a different result.")
}

func incrementCounter(counter *int) {
	*counter++
}
func decrementCounter(counter *int) {
	*counter--
}
