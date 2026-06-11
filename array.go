package main

import "fmt"

func main() {
	str := [3]string{}
	str[0] = "Hello"
	str[1] = "World"
	str[2] = "!"
	fmt.Println(str)
	str2 := [...]string{"Hello", "World", "!"}
	fmt.Println(str2)
}