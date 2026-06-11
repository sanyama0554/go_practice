package main

import "fmt"

func main() {
	a1 := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
	}
	
	fmt.Println(a1["two"])
	a1["four"] = 4
	fmt.Println(a1)
	delete(a1, "two")
	fmt.Println(a1)
	_, ok := a1["two"]	// _ is used to ignore the value, we only care about whether the key exists or not
	if ok {
		fmt.Println("two is in the map")
	} else {
		fmt.Println("two is not in the map")
	}

	for key, value := range a1 {
		fmt.Printf("%s: %d\n", key, value)
	}
}