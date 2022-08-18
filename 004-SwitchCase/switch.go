package main

import "fmt"

func main() {
	a := -100

	switch {
	case a > 0:
		if a%2 == 0 {
			break
		}
		fmt.Println("Odd positive value received")
	case a < 0:
		fmt.Println("Negative value received")
		fallthrough
	case a > -10:
		fmt.Println("-10")
	case a < -20:
		fmt.Println("-20")
	default:
		fmt.Println("Default value handling")
	}
	fmt.Print("Введите год рождения: ")
	var year int
	n, err := fmt.Scanln(&year)

	switch {
	case err != nil || n < 1 || year < 1946:
		fmt.Println("Некорректный ввод!")
	case year <= 1964:
		fmt.Println("Привет, бумер!")
	case year <= 1980:
		fmt.Println("Привет, представитель X!")
	case year <= 1996:
		fmt.Println("Привет, миллениал!")
	case year <= 2012:
		fmt.Println("Привет, зумер!")
	default:
		fmt.Println("Привет, альфа!")
	}
}
