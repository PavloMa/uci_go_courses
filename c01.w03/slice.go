/* 
Write a program which prompts the user to enter integers and stores the integers 
in a sorted slice. The program should be written as a loop. Before entering the loop, 
the program should create an empty integer slice of size (length) 3. 
During each pass through the loop, the program prompts the user to enter an integer 
to be added to the slice. The program adds the integer to the slice, sorts the slice, 
and prints the contents of the slice in sorted order. 
The slice must grow in size to accommodate any number of integers which the user 
decides to enter. The program should only quit (exiting the loop) when the user 
enters the character ‘X’ instead of an integer.
*/

package main

import (
	"fmt"
	"sort"
	"encoding/json"
	"strconv"
)

func main() {

	numbers := make([]int, 0, 3)

	for i := 1; ; i++ {
		fmt.Print("Enter an integer to be appended to the slice or 'X'/'x' to quit:")

		var line string  
		fmt.Scan(&line)
		if line == "X" || line == "x" {
			fmt.Println("Bye!")
			return
		}

		integer, err := strconv.Atoi(line)
	    if err != nil {
	        fmt.Println("Wrong input, try again")
	        continue
	    }
	    
	    numbers = append(numbers, integer)
		sort.Ints(numbers)

    	formatted, _ := json.Marshal(numbers)
		fmt.Printf("%v%v\n", "Sorted slice: ", string(formatted)) // ["1", "2", "A", "B", "Hack\"er"]
    	//fmt.Println("[", numbers, "]")
	}

    // fmt.Println("Bye!")
	 
}

