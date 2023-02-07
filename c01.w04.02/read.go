/*Write a program which reads information from a file and represents it 
in a slice of structs. Assume that there is a text file which contains 
a series of names. Each line of the text file has a first name and a last name, 
in that order, separated by a single space on the line. 

Your program will define a name struct which has two fields, fname for the first 
name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. 
Your program will successively read each line of the text 
file and create a struct which contains the first and last names found in the file. 
Each struct created will be added to a slice, and after all lines have 
been read from the file, your program will have a slice containing one struct 
for each line in the file. After reading all lines from the file, your program 
should iterate through your slice of structs and print the first and last names 
found in each struct.

Submit your source code for the program, “read.go”.*/

package main

import (
	"fmt"
	"bufio"
	"os"
	//"io"
	//"strings"
)

type Person struct {
	fName string
	lName string
}

func main() {

    // a scanner to read input with spaces
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a file name: ")
	scanner.Scan() 
	fileName := scanner.Text()

	scanner = nil //close this scanner
	
	file, err := os.Open(fileName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "os.Open error: %v\n", err)
        return
    }

    // fmt.Println() //empty string

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    var persons []Person = make([]Person, 0, 3)

    for fileScanner.Scan() {
    	line := fileScanner.Text()
    	// fmt.Println(line)

		var (
		    firstName string
		    lastName string
		)
		
		fmt.Sscanln(line, &firstName, &lastName)
        
    	person := Person{truncateString(firstName, 20), truncateString(lastName, 20)}

    	persons = append(persons, person)
    }
  
    file.Close()

    // fmt.Println("Items: ", len(persons))

	for i, element := range persons {
        fmt.Printf("%v %s %s\n", i, element.fName, element.lName)
    }

}


func truncateString(input string, length int) string {
	if len(input) > length {
		return string(input[0:length])
	} else {
		return input
	}
}

	/*The fmt.Fscanf() function in Go language scans the specified text, 
	read from r and then stores the successive space-separated values into 
	successive arguments as determined by the format. Here newlines in the 
	input must match newlines in the format.*/

	//num, err = Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
	

        //reader := strings.NewReader(line)

		//fmt.Fscanln(reader, &firstName, &lastName) // from the string, not from file
  		// _, err
        /*if err == io.EOF {
            break
        } */       

		/*		if err == io.EOF {
			break
		}

		if err != nil {
	        fmt.Fprintf(os.Stderr, "Fscanln error: %v\n", err)
	        return
	    }

    	person := Person{truncateString(firstName, 20), truncateString(lastName, 20)}

    	persons = append(persons, person)

    } */
