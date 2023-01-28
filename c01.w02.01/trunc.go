// Truncating

package main

import (
	"fmt"
	"log"
	//"https://github.com/rung/go-safecast"
)

func main() {
	var floating float64
	var integer int64

	// Calling Scanf() function for
	// reading the input from the standard input
	_, err := fmt.Scanf("%g", &floating)
    if err != nil {
        log.Fatal(err)
        fmt.Println("Erroneous input")
        return 
    }

    //unsafe typecast:
	integer = int64(floating)

	//safe typecast:

	/*integer, err := safecast.Int64(floating) // convert int to int64 in a safe way
	if err != nil {
        log.Fatal(err)
        fmt.Println("Cannot safely convert float to int")
		return err
	}*/

	fmt.Println(integer)

}

