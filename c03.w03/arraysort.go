/*
Write a program to sort an array of integers. 
The program should partition the array into 4 parts, 
each of which is sorted by a different goroutine. 
Each partition should be of approximately equal size. 
Then the main goroutine should merge the 4 sorted 
subarrays into one large sorted array. 

The program should prompt the user to input 
a series of integers. Each goroutine which sorts ¼ o
f the array should print the subarray that it will sort. 
When sorting is complete, the main goroutine should print the entire sorted list.
*/

/*ATTENTION
There is problem printing from concurrent goroutines. Printed text can be interleaved*/

package main

import (
	"fmt"
	//"encoding/json"
	//"strconv"
	"math"
	"sync"
)

func main() {

	fmt.Println("Enter a sequence of enter-separated integers and press extra enter:")
	ints := InputInts() // PRODUCTION
	// var ints = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1} // TESTING
	fmt.Print("Unsorted array: ")
	// fmt.Println()
	PrintSlice(ints)
	fmt.Println()

	length := len(ints) 
	// DEBUG fmt.Println("Len: ", length)
	if length < 1 {
		return
	}

	const sections int = 4

	if length < sections { // if the array is short, just sort it all at once
		BubbleSort(ints)
	} else {
		ints = SectionSort(ints, sections)
	}
	
	fmt.Print("Sorted array: ")
	PrintSlice(ints)

}

func InputInts() []int {
    var num int
    ints := make([]int, 0, 10)

    for {
        qty, _ := fmt.Scanf("%d", &num)
        if qty == 1 {
	        ints = append(ints, num)
	        // DEBUG PrintSlice(ints)
	    } else {
            break // requirements do not prescribe behaviour given erroneous input
        }
    }
    return ints
}	


func SectionSort(ints []int, sections int) []int {

	if !IsPositivePowerOfTwo(sections) {
		fmt.Println("Number of array slices must be power of 2")
		return nil
	}

	length := len(ints)
	streak := int(math.Round(float64(length)/float64(sections))) // FP division with rounding. Explained below
	
	indices := make([]int, sections + 1, sections + 1) // start indices of the 4 array sections + a fake index of the next for symmetry
	for i := 0; i < sections; i++ {
		indices[i] = streak*i
	}
	indices[sections] = length // a fake index ouside the range, needed to calculate the end of the fourth streak
	// DEBUG fmt.Printf("Ints length: %d, Streak: %d, Indices length: %d\n", length, streak, len(indices))
	// DEBUG fmt.Print("Breakdown indices: ")
	// DEBUG PrintSlice(indices)
	// DEBUG fmt.Println("")
	// Examples (integer division - quotient):
	// streak := length/sections // integer division
	// length = 4. Indices 0, 1, 2, 3. streak = 1. Indices: 1*0=0, 1*1=1, 1*2=2, 1*3=3
	// length = 5. Indices 0, 1, 2, 3, 4. streak = 1. Indices: 1*0=0, 1*1=1, 1*2=2, 1*3=3
	// length = 7. Indices 0, 1, 2, 3, 4, 5, 6. streak = 1. Indices: 1*0=0, 1*1=1, 1*2=2, 1*3=3
	// length = 8. Indices 0, 1, 2, 3, 4, 5, 6, 7. streak = 2. Indices: 2*0=0, 2*1=2, 2*2=4, 2*3=6
	// so, the last subarray is often longer than the first three
	// a more balanced way would be to use FP arithmetic and rounding (not perfect either)
	// but it makes calculating indices for merge complicated:
	// streak := int(math.Round(float64(length)/float64(sections)))
	// instead of streak := length/sections

    var wg sync.WaitGroup
    wg.Add(sections)

    var m sync.Mutex

	for i := 0; i < sections; i++ {

	    go func(wg *sync.WaitGroup, m *sync.Mutex, numbers []int, thread int) {
	        
	        m.Lock() // Ensure the arrays are printed without interleaving
	        fmt.Printf("Unsorted subarray %d: ", thread)
	        PrintSlice(numbers)
	        fmt.Println(" ")
	        m.Unlock()

	        BubbleSort(numbers)
	        //fmt.Println("")
	        (*wg).Done()
	    } (&wg, &m, ints[indices[i]:indices[i+1]], i)
	}

    wg.Wait()
	return Merge(ints, indices)
}

func PrintSlice(numbers []int) { 
	// TODO: better making JSON. See below
	fmt.Print("[")
	length := len(numbers)
	for i := 0; i < length; i++ {
		fmt.Print(numbers[i], ", ")
	} 
	fmt.Print("\b\b]") // erasing the last comma and space
	// formatted, _ := json.Marshal(numbers)
	//fmt.Printf("%v%v\n", "Sorted slice: ", string(formatted)) // ["1", "2", "A", "B", "Hack\"er"]
}

func BubbleSort(numbers []int) { // sort a slice of an array from start to next excluding
	
	length := len(numbers) 

	for j := length - 1; j > 0; j-- { //elements at the end are already sorted
		for i := 0; i < j; i++ { //no need to compare the last single element
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


func Merge(numbers []int, indices []int) []int { // merge sorted sections of an array 

	// idea from the merge sort:
	// [^2, 4] [^1, 3] > [1]; [2, ^4] [^2, 3] > [1, 2]; [2, ^4] [1 ^3] ...

	sections := len(indices) - 1 //the last index is fake, it shows the first element outside the range
	if sections < 2 { //nothing to merge. Just return the input
		return numbers
	} else if sections == 2 { //merge two sections
		a := numbers[indices[0]:indices[1]]
		b := numbers[indices[1]:indices[2]]
		i := 0
		j := 0
		result := []int{}
	    for i < len(a) && j < len(b) { // process until reaching the end of any of the two arrays
	        if a[i] < b[j] {
	            result = append(result, a[i])
	            i++
	        } else {
	            result = append(result, b[j])
	            j++
	        }
	    }
	    for ; i < len(a); i++ { // process the tail of the first array, if any
	        result = append(result, a[i])
	    }
	    for ; j < len(b); j++ { // process the tail of the second array, if any
	        result = append(result, b[j])
	    }
	    return result
	} else { // sections > 2, i.e. 4 or 8 or ... 
		mid := sections / 2 // sec = 4, mid =2, Indices: 0, 1, 2, and 2, 3, 4. Two is overlap
		// It is a merge of merge here, with new combined indices array to be passed
		newIndices := []int{indices[0], indices[mid], indices[len(indices)-1]}
		// DEBUG fmt.Print("Breakdown indices: ")
		// DEBUG PrintSlice(newIndices)
		// DEBUG fmt.Println("")
		result := append(
						append([]int{}, Merge(numbers, indices[:mid+1])...), 
			       		Merge(numbers, indices[mid:])...
			       	    )
		return Merge(result, newIndices)
	}

}

func IsPositivePowerOfTwo(x int) bool {
    return (x > 0) && ((x & (x - 1)) == 0);
}