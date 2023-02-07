/* 
Write a program which allows the user to get information about a predefined set of animals. 
Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. 
The user can issue a request to find out one of three things about an animal: 
1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when 
it speaks. The following table contains the three animals and their associated data 
which should be hard-coded into your program.

Animal Food eaten	Locomotion method 	Spoken sound
“cow”	grass	walk	moo
“bird”	worms	fly		peep
“snake”	mice	slither	hsss 

Your program should present the user with a prompt, “>”, to indicate that the user can 
type a request. Your program accepts one request at a time from the user, prints out 
the answer to the request, and prints out a new prompt. Your program should continue 
in this loop forever. Every request from the user must be a single line containing 2 strings. 
The first string is the name of an animal, either “cow”, “bird”, or “snake”. 
The second string is the name of the information requested about the animal, 
either “eat”, “move”, or “speak”. Your program should process each request by printing 
out the requested data.

You will need a data structure to hold the information about each animal. Make a type 
called Animal which is a struct containing three fields:food, locomotion, and noise, 
all of which are strings. Make three methods called Eat(), Move(), and Speak(). 
The receiver type of all of your methods should be your Animal type. The Eat() method 
should print the animal’s food, the Move() method should print the animal’s locomotion, 
and the Speak() method should print the animal’s spoken sound. Your program should call 
the appropriate method when the user makes a request.
*/

package main

import (
	"fmt"
)

type Animal struct {
	name string
	food string
	locomotion string
	noise string
}

func (am Animal) Name () {
	fmt.Println("Animal: ", am.name)
}

func (am Animal) Eat () {
	am.Name()
	fmt.Println("Food: ", am.food)
}

func (am Animal) Move () {
	am.Name()
	fmt.Println("Movement: ", am.locomotion)
}

func (am Animal) Speak () {
	am.Name()
	fmt.Println("Sound: ", am.noise)
}

func main() {

	cow := Animal{"cow", "grass", "walk", "moo"}
	bird := Animal{"bird", "worms", "fly", "peep"}
	snake := Animal{"snake", "mice", "slither", "hsss"}
	var animals = map[string]Animal {
        "cow": cow,
        "bird": bird,   
        "snake": snake,// last comma is a must
    }

	fmt.Println("Enter one animal name (cow, bird or snake) \nand a requested information (food, movement or sound) separated by space\nor 'x x' to quit:")
	
	for {
		fmt.Print(">")

		var animal string
		var info string

		fmt.Scanf("%s %s", &animal, &info)

		var am Animal = animals[animal]
		switch animal {
		  	case "cow","bird", "snake":
		  		am = animals[animal]
		    case "x":

			default:
			    fmt.Println("Animal name is incorrect")
			    continue
		}		

		switch info {
		  	case "food":
		  		am.Eat()
		  	case "movement":
		    	am.Move()
		  	case "sound":
		    	am.Speak()
		    case "x":
		    	return
			default:
			    fmt.Println("Info name is incorrect")
			    continue
		}	

	}

}
