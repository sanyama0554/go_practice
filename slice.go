package main

import "fmt"

func main() {
	str := []string{}
	str = append(str, "Hello")
	str = append(str, "World")
	str = append(str, "!")
	fmt.Println(str)
	fmt.Println(len(str), cap(str))
}