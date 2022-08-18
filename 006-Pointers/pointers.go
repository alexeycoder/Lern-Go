package main

import "fmt"

func main() {

	var someVariable int = 5
	pointer := &someVariable

	fmt.Println(someVariable, pointer) //a=5 p=0xc0000b2008

	a := 1
	p := &a
	b := &p

	*p = 3
	**b = 5

	a += 4 + *p + **b

	fmt.Printf("%d\n", *p)
}
