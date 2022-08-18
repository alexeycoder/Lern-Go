package main

import "fmt"

func main() {
	fmt.Println("Break label:")

outerLoopLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			break outerLoopLabel
		}
	}
	fmt.Println("End")

	fmt.Println("\nContinue label:")

outerLoopSecondLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			continue outerLoopSecondLabel
		}
	}
	fmt.Println("End")

	fmt.Println("\nEvens with decades:")
	group := 0
	for i := 0; i < 20; i++ {
		switch {
		case i%2 == 0:
			if i%10 == 0 {
				group++
				break
			}
			fmt.Printf("%02d: %d\n", group, i)
		default:
		}
	}

	fmt.Println("\nFizzBuzz:")
	for i := 1; i <= 100; i++ {
		isMultipleOfThree := i%3 == 0
		isMultipleOfFive := i%5 == 0
		switch {
		case isMultipleOfThree && isMultipleOfFive:
			fmt.Print("FizzBuzz")
		case isMultipleOfThree:
			fmt.Print("Fizz")
		case isMultipleOfFive:
			fmt.Print("Buzz")
		default:
			fmt.Print(i)
		}
		fmt.Print(" ")
	}
	fmt.Println()
}
