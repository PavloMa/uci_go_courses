/* 
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. 
The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, 
but it should swap the contents of the slice in position i with the contents in 
position i+1.
*/

package main

import (
	"fmt"
	// "sort"
	"encoding/json"
	"strconv"
)

func main() {

	numbers := make([]int, 0, 10)

	fmt.Println("Enter a sequence of up to ten integer to be appended to the slice\nor 'X'/'x' to quit:")
	for i := 1; i <= 10; i++ {
		//fmt.Print("Enter an integer to be appended to the slice or 'X'/'x' to quit:")

		var line string  
		fmt.Scan(&line)
		if line == "X" || line == "x" {
			//fmt.Println("Bye!")
			break //stop getting ints
		}

		integer, err := strconv.Atoi(line)
	    if err != nil {
	        fmt.Println("Wrong integer formqt, try again")
	        continue
	    }
	    
	    numbers = append(numbers, integer)
	}

	BubbleSort(numbers)

	formatted, _ := json.Marshal(numbers)
	fmt.Printf("%v%v\n", "Sorted slice: ", string(formatted)) // ["1", "2", "A", "B", "Hack\"er"]

}

func BubbleSort(numbers []int) {
	
	length := len(numbers)

	for j := length - 1; j > 0; j-- { //elements at the end are already sorted
		for i := 0; i < j; i++ { //no need to compare the last element
			if (numbers[i] > numbers[i + 1]) {
				Swap(numbers, i)
			}
		}
	}
}

func Swap(numbers []int, firstIndex int) {
	i := firstIndex
	if (i + 1 < len(numbers)) { // check if index is in range
		tmp := numbers[i + 1]
		numbers[i + 1] = numbers[i]
		numbers[i] = tmp
	}
}


