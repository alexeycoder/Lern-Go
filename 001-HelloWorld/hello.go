package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var a float32 = 0.56
	var b float32 = 1.74
	fmt.Println(a / b)

	var check bool = true

	if check {
		var s string
		s = "abc"
		fmt.Println(string(s[2]))
		println(len(s))
	}

	var i int
	var i32 int32
	println(i == int(i32))
}
