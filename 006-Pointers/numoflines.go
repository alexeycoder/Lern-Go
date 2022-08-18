package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Получаем читателя пользовательского ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	cnt := 0
	for {
		fmt.Print("-> ")
		// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
		str, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&cnt)

		fmt.Printf("User input %d lines\n", cnt)

		if str == "end\n" {
			break
		}
	}

	tmp := 34567
	po := &tmp

	fmt.Println(*po)
}

func f(cnt *int) {
	*cnt++
}