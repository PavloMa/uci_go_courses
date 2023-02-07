package main

import (
	"fmt"
)

func main() {
	ary := getInputs()
	fmt.Print("\n")

	a1, a2, a3, a4 := divideArray(ary)
	fmt.Print("Divided arrays:\n")
	printAry("a1 = ", a1)
	printAry("a2 = ", a2)
	printAry("a3 = ", a3)
	printAry("a4 = ", a4)
	fmt.Print("\n")

	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	c4 := make(chan int)

	go cBubbleSort(c1, a1)
	go cBubbleSort(c2, a2)
	go cBubbleSort(c3, a3)
	go cBubbleSort(c4, a4)

	ret := 0
	ret = <- c1
	ret = <- c2
	ret = <- c3
	ret = <- c4
	ret = ret + 1

	fmt.Print("Divided and sorted arrays:\n")
	printAry("a1 = ", a1)
	printAry("a2 = ", a2)
	printAry("a3 = ", a3)
	printAry("a4 = ", a4)
	fmt.Print("\n")

	aa := []int{}
	aa = append(a1, a2...)
	aa = append(aa, a3...)
	aa = append(aa, a4...)
	BubbleSort(aa)
	
	fmt.Print("Combined and sorted results:\n")
	printAry("", aa)
}

func getInputs() []int {
	var x int
	ary := make([]int, 0)

	fmt.Println("Enter a list of integers. Type '.' to finish entering numbers.: ")
	for {
		n, _ := fmt.Scan(&x)
		if n == 0 {
			break
		}
		for j := 0; j < n; j++ {
			ary = append(ary, x)
		}
	}

	return ary
}

// swap the contents at i and i+i
func Swap(sli []int, i int) {
	t0 := sli[i]
	t1 := sli[i+1]
	sli[i] = t1
	sli[i+1] = t0
}

func cBubbleSort(c chan int, sli []int) {
	BubbleSort(sli)
	c <- 1
}

func BubbleSort(sli []int) {
	for {
		swapped := false
		for j := 0; j < len(sli)-1; j++ {
			t0 := sli[j]
			t1 := sli[j+1]
			if t0 > t1 {
				Swap(sli, j)
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}

func printAry(s string, ary []int) {
	fmt.Printf("%s", s)
	for i := 0; i < len(ary); i++ {
		fmt.Printf("%d ", ary[i])
	}
	fmt.Println("")
}

func divideArray(ary []int) ([]int, []int, []int, []int) {
	length := len(ary)
	len4 := length / 4
	a1 := ary[0:len4]
	a2 := ary[(1 * len4): 2 * len4]
	a3 := ary[(2 * len4): 3 * len4]
	a4 := ary[(3 * len4): length]

	a1x := make([]int, len(a1))
	copy(a1x, a1)
	a2x := make([]int, len(a2))
	copy(a2x, a2)
	a3x := make([]int, len(a3))
	copy(a3x, a3)
	a4x := make([]int, len(a4))
	copy(a4x, a4)

	return a1x, a2x, a3x, a4x
}
