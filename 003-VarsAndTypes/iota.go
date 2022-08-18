package main

import "fmt"

const (
	_ = iota * 10
	ten
	hundred
	thousand
)

const (
	hello  = "Hello, world!" // iota = 0 here
	numOne = 1               // iota = 1 here

	black = iota
	gray
)

const (
	one = iota*2 + 1
	three
	five
	seven
	nine
	eleven
)

const height uint8 = iota + 1

func main() {
	fmt.Println(ten, hundred, thousand)
	fmt.Println(hello, one, black, gray)
	fmt.Println(one, three, five, seven, nine, eleven)
	fmt.Println(height)
}
