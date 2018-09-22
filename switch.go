package main

import "fmt"
import "time"

func main() {

	// Normal Switch looks like this
	i := 3
	fmt.Println("Write", i, "as")
	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
	}

	// With Default case
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's weekend bitches!!!")
		default:
			fmt.Println("It's Weekday :')")
	}

	// Without condition it behaves like alternate of if/else
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("It's before Noon")
		default:
			fmt.Println("It's after noon")
	}

	whatamI := func (i interface{}) {
		switch t := i.(type) {
			case bool:
				fmt.Println("I'm a Bool")
			case int:
				fmt.Println("I'm int")
			default:
				fmt.Println("Don't know what the fuck I'm", t)
		}
	}

	whatamI(true)
	whatamI(1)
	whatamI("hey")
}