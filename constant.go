package main

import "fmt"

func main() {
	const a1 = 1
	fmt.Println(a1)
	a1 = 2 // error: cannot assign to a1
}