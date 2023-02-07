package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func requestUserEntry(constructionText string) []int {
	var userNumbers = []int{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(constructionText)
	scanner.Scan()
	userInput := scanner.Text()
	numbers_split := strings.Split(userInput, " ")
	for _, numberStr := range numbers_split {
		convVal, err := strconv.Atoi(numberStr)
		if err == nil {
			userNumbers = append(userNumbers, convVal)
		}
	}
	return userNumbers
}

func BubbleSort(numbers []int, channel chan []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i)
				swapped = true
			}
		}
	}
	//fmt.Printf("Sorting following list [%v]\n", numbers)
	channel <- numbers
}

func Swap(numbers []int, pos int) {
	pos_next := pos + 1
	curr := numbers[pos]
	numbers[pos] = numbers[pos_next]
	numbers[pos_next] = curr
}

func main() {
	channel := make(chan []int)
	requestText := "Enter your list of integers to be sorted seperated by whitespace and enter when you are done\n"
	listToSort := requestUserEntry(requestText)
	listLength := len(listToSort)
	sliceLength := int(math.Ceil(float64(listLength) / 4))
	for i := 0; i < listLength; i += sliceLength {
		if i+sliceLength < listLength {
			//fmt.Printf("%d - %d\n", i, i+sliceLength)
			go BubbleSort(listToSort[i:i+sliceLength], channel)
		} else {
			//fmt.Printf("%d - until end %d", i, listLength)
			go BubbleSort(listToSort[i:], channel)
		}
	}

	sp1, sp2, sp3, sp4 := <-channel, <-channel, <-channel, <-channel
	completeList := append(sp1, sp2...)
	completeList = append(completeList, sp3...)
	completeList = append(completeList, sp4...)
	sort.Ints(completeList)

	fmt.Printf("Partially Sorted: %v, %v, %v, %v\n", sp1, sp2, sp3, sp4)
	fmt.Printf("Sorted complete List: %v", completeList)
}
