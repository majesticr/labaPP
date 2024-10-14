package main

import (
	"fmt"
)

func main() {
	var n int
	var sum int

	fmt.Println("Введите числа (0 - для выхода):")
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		sum += n
	}

	// Вывод суммы
	fmt.Printf("Sum: %d\n", sum)
}
