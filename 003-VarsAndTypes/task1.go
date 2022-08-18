package main

import "fmt"

const (
	i         = 10
	f         = 1.5
	i64 int64 = 88
)

func main() {
	ver := "v0.0.1"
	id := 0
	pi := 3.1415
	fmt.Println("ver =", ver, "id =", id, "pi =", pi)

	var v = 45

	fmt.Println(i + f)
	fmt.Println(i + i64)
	//fmt.Println(f+ untyped float(i64))
	fmt.Println(i64 + int64(v))
	fmt.Println(i + v)
	//fmt.Println(f+untyped float(v))

}
