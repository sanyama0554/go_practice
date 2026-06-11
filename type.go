package main

func main() {
	type UtcTime string
	type JstTime string

	var t1 UtcTime = "2024-06-01T12:00:00Z"
	var t2 JstTime = "2024-06-01T21:00:00+09:00"

	t1 = t2 // error: cannot use t2 (type JstTime) as type UtcTime in assignment
}