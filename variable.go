package main
import "fmt"

func main() {
	var a1 int = 1
	var a2 = 2
	var (
		a3 int = 3
	)
	a4 := 4
	name, age := "Yuma", 25
	
	fmt.Println(a1, a2, a3, a4)
	fmt.Println(name, age)
}