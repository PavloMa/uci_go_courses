package main

import (
	"fmt"
	"strconv"
)

func BubbleSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

func Swap(numbers []int, index int) {
	temp := numbers[index]
	numbers[index] = numbers[index+1]
	numbers[index+1] = temp
}

func main() {
	fmt.Println("Type a sequence of up to 10 integers!\nAdd 'x' when you want to stop!\nExample: '4 3 2 5 x', '1 2 3 4 5 6 7 8 9 10'")
	numbers := make([]int, 0, 1)
	var ch string
	for i := 0; i < 10; i++ {
		fmt.Scan(&ch)
		if ch == "x" {
			break
		}
		num, _ := strconv.Atoi(ch)
		numbers = append(numbers, num)
	}
	if len(numbers) == 0 {
		fmt.Println("\nMust enter at least one number to sort!")
		return
	}
	BubbleSort(numbers)
	fmt.Println("\nYour numbers in sorted order from least to greatest...")
	for _, v := range numbers {
		fmt.Println(v)
	}
}
