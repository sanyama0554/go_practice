package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	colors := []string{"red", "green", "blue"}
	for _, color := range colors {
		fmt.Println(color)
	}
}