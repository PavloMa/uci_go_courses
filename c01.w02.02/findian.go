/*
Input from console and search a string for the pattern: i*a*n

Write a program which prompts the user to enter a string. 
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. 
The program should print “Found!” if the entered string starts with the character ‘i’, 
ends with the character ‘n’, and contains the character ‘a’. 
The program should print “Not Found!” otherwise. 
The program should not be case-sensitive, 
so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, 
“ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. 
The program should print “Not Found!” for the following strings, 
“ihhhhhn”, “ina”, “xian”. 
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
)

var input string

func main() {
	fmt.Println("Enter a string:")

        // read input with spaces
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() 
	input := scanner.Text()

	if (checkPattern(input)) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}

func checkPattern(s string) bool {
	const pattern = `(?i)^i.*a.*n$` // (?i) - case insensitive
	result, _ := regexp.MatchString(pattern, s)
	return result
}

