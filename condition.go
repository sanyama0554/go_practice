package main

import "fmt"

func main() {
	a1 := 1
	a2 := 2

	if a1 < a2 {
		fmt.Printf("%d is less than %d\n", a1, a2)
	} else if a1 > a2 {
		fmt.Printf("%d is greater than %d\n", a1, a2)
	} else {
		fmt.Printf("%d is equal to %d\n", a1, a2)
	}

	switch a1 {
	case 1:
		// fmt.Println("a1 is 1")
		fallthrough	// fallthrough is used to execute the next case even if the condition is not met
	case 2:
		fmt.Println("a1 is 2")
	default:
		fmt.Println("a1 is neither 1 nor 2")
	}
}