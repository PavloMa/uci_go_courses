/*Assignment: Write two goroutines which have a race condition 
when executed concurrently. Explain what the race condition is and how it can occur.*/

/*Explanation: two gourutines, which increment and decrement the same counter, are started
concurrently 50 times each. While they include non-atomic operations their execution may 
interleave leading to race conditions.*/

/*Testing for race conditions: run the program as follows:
go run -race race.go
The utility provides info about the detected race condition, in particular:
Found 1 data race(s)
exit status 66
*/

package main

import (
    "fmt"
    //"runtime"
    // "sync"
    "time"
)

var counter int32

func increment() {
    time.Sleep(50 * time.Millisecond)
    counter++
    fmt.Println("increment:", counter)
}

func decrement() {
    time.Sleep(50 * time.Millisecond)
    counter--
    fmt.Println("decrement:", counter)
}

func main() {

    counter = 0

    for i := 0; i < 50; i++ {
        go increment() 
        go decrement()
    }
    time.Sleep(time.Second * 4)
}