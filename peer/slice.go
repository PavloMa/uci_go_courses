package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func main() {
	var intSlice = make([]int, 0, 3)
	var userInput string
	
	for strings.Compare(userInput, "X") != 0 {
		fmt.Println("Enter an integer:")
		fmt.Scan(&userInput)
		i, err := strconv.Atoi(userInput)

		if err == nil {
			intSlice = append(intSlice, i)
			sort.Ints(intSlice)
		}
		
		if len(intSlice) > 0 {
			sliceString := strings.Trim(strings.Join(strings.Split(fmt.Sprint(intSlice), " "), ", "), "[]")
			fmt.Println(sliceString)
		}
	}
}
