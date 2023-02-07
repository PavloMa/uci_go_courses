/*Write a program which prompts the user to first enter a name, 
and then enter an address. Your program should create a map and add the name 
and address to the map using the keys “name” and “address”, respectively. 
Your program should use Marshal() to create a JSON object from the map, 
and then your program should print the JSON object.*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)

type Person struct {
    Name string
    Address string
}

func main() {

    // a scanner is necessary to read input with spaces
	scanner := bufio.NewScanner(os.Stdin)


	fmt.Print("Enter a name: ")
	scanner.Scan() 
	name := scanner.Text()

	fmt.Print("Enter an address: ")
	scanner.Scan() 
	address := scanner.Text()

    // assure proper indentation
    bytes, err := json.MarshalIndent(Person{Name: name, Address: address}, "", "    "); // 4 spaces
    if err != nil {
        fmt.Println("Error marshalling json")
        return
    }
    fmt.Println(string(bytes)) 
}

