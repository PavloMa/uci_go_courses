/* 
Write a program which allows the user to create a set of animals 
and to get information about those animals. 
Each animal has a name and can be either a cow, bird, or snake. 
With each command, the user can either create a new animal of one of the three types, 
or the user can request information about an animal that he/she has already created. 
Each animal has a unique name, defined by the user. 
Note that the user can define animals of a chosen type, 
but the types of animals are restricted to either cow, bird, or snake. 
The following table contains the three types of animals and their associated data.

animal 	Food  	Locomtion 	Sound
cow 	grass	walk	moo
bird 	worms 	fly 	peep
snake 	mice 	slither hsss 

Your program should present the user with a prompt, “>”, 
to indicate that the user can type a request. 
Your program should accept one command at a time from the user, 
print out a response, and print out a new prompt on a new line. 
Your program should continue in this loop forever. 
Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. 
The first string is “newanimal”. 
The second string is an arbitrary string which will be the name of the new animal. 
The third string is the type of the new animal, either “cow”, “bird”, or “snake”.  
Your program should process each newanimal command by creating the new animal 
and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. 
The first string is “query”. 
The second string is the name of the animal. 
The third string is the name of the information requested about the animal, 
either “eat”, “move”, or “speak”. 
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. 
Specifically, the Animal interface should contain the methods 
Eat(), Move(), and Speak(), which take no arguments and return no values. 
The Eat() method should print the animal’s food, the Move() method should print 
the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. 
Define three types Cow, Bird, and Snake. For each of these three types, define methods 
Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the 
Animal interface. 
When the user creates an animal, create an object of the appropriate type. 
Your program should call the appropriate method when the user issues a query command.
*/

package main

import (
	"fmt"
)

//Animal

type Animal interface {
	Eat()
	Move()
	Speak()
	Name() 
}

// Cow

type Cow struct {
	//species string 
	name string
	food string
	locomotion string
	noise string
}

func NewCow(name string) *Cow {
  return &Cow{name, "grass", "walk", "moo"}
}

func (am Cow) Name () {
	fmt.Println("Name: ", am.name)
}

func (am Cow) Eat () {
	fmt.Println("Food: ", am.food)
}

func (am Cow) Move () {
	fmt.Println("Movement: ", am.locomotion)
}

func (am Cow) Speak () {
	fmt.Println("Sound: ", am.noise)
}

// Bird

type Bird struct {
	//species string 
	name string
	food string
	locomotion string
	noise string
}

func NewBird(name string) *Bird {
  return &Bird{name, "worms", "fly", "peep"}
}

func (am Bird) Name () {
	fmt.Println("Name: ", am.name)
}

func (am Bird) Eat () {
	fmt.Println("Food: ", am.food)
}

func (am Bird) Move () {
	fmt.Println("Movement: ", am.locomotion)
}

func (am Bird) Speak () {
	fmt.Println("Sound: ", am.noise)
}

// Snake

type Snake struct {
	//species string 
	name string
	food string
	locomotion string
	noise string
}

func NewSnake(name string) *Snake {
  return &Snake{name, "mice", "slither", "hsss"}
}

func (am Snake) Name () {
	fmt.Println("Name: ", am.name)
}

func (am Snake) Eat () {
	fmt.Println("Food: ", am.food)
}

func (am Snake) Move () {
	fmt.Println("Movement: ", am.locomotion)
}

func (am Snake) Speak () {
	fmt.Println("Sound: ", am.noise)
}


func createAnimal(name string, species string, animals map[string]Animal) { //map is actually a reference
	var am Animal
	switch species {
	  	case "cow":
	  		am = NewCow(name)
	  	case "bird":
	  		am = NewBird(name)
	  	case "snake":
	  		am = NewSnake(name)
	  	default:
	  		fmt.Println("The %s is a wrong species. Must be one of (cow|bird|snake)\n", species)
	  		return
	}
	_, exists := animals[name] // not "am", otherwise it overwrites it to nil
    if (exists) {
        fmt.Printf("Animal named %s already exists. Must be a new name\n", name) 
        return
    } else {  
	    fmt.Println("Created it!")
		animals[name] = am // add the new animal to the map
	}
}

func queryInfo(name string, info string, animals map[string]Animal) {
	am, exists := animals[name]
    if (exists) {
		switch info {
		  	case "food":
		  		am.Eat()
		  	case "movement":
		    	am.Move()
		  	case "sound":
		    	am.Speak()
			default:
			    fmt.Printf("Requested info %s is unavailable. Must be one of (food|movement|sound)\n", info)
		}
	} else {
        fmt.Printf("Animal named %s is not yet created.\n", name) 
    }   
}

func wrongCommand(name string) {
    fmt.Printf("Command %s is incorrect.\n", name)
}

func main() {

	var animals = make(map[string]Animal) 

	for {
		fmt.Println("\nEnter one of the two commands (newanimal|query)",
					"in the following format (three words each):\n",
					"newanimal animal_name species(cow|bird|snake)\n",
					"query animal_name information(food|movement|sound)\n",
					"or press 'Ctrl/Cmd-C' to quit:")
		fmt.Print(">")

		var command string
		var name string
		var argument string

		fmt.Scanf("%s %s %s", &command, &name, &argument)
		switch command { // choose a command processor
		  	case "newanimal":
		  		species := argument
		  		createAnimal(name, species, animals)
		  		// DEBUG fmt.Println("Map is:", animals)
				// DEBUG fmt.Println("Length of map :", len(animals))
		  	case "query":
		  		info := argument
		  		queryInfo(name, info, animals)
			default:
				wrongCommand(command)
		}		
	}

}
