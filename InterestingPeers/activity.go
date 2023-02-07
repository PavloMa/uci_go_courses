/*
You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))

Submit your Go program source code.
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*(t*t) + vo*t + so
	}
}

func ArgsPrompt() map[string]float64 {
	fmt.Print("Please enter acceleration, initial velocity, and initial displacement as comma separated parameters (a, b, c).\n")
	var input string
	fmt.Scan(&input)
	args := strings.Split(input, ",")
	if len(args) < 3 {
		fmt.Println("Please enter 3 parameters.")
		return nil
	}
	a, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("Please make sure all 3 parameters are valid numbers.")
		return nil
	}
	vo, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("Please make sure all 3 parameters are valid numbers.")
		return nil
	}
	so, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Println("Please make sure all 3 parameters are valid numbers.")
		return nil
	}

	return map[string]float64{
		"a":  a,
		"vo": vo,
		"so": so,
	}
}

func main() {
	argsMap := ArgsPrompt()
	fn := GenDisplaceFn(argsMap["a"], argsMap["vo"], argsMap["so"])
	if argsMap != nil {
		var input string
		for {
			fmt.Print("Please enter a value for time ('X' to end):")
			fmt.Scan(&input)
			if input == "X" {
				break
			}
			t, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Please make sure it is a number.")
			} else {
				fmt.Println("Displacement: ", fn(t))
			}

		}
	}

}
