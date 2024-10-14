package main

import (
	"fmt"
)

// Функция для генерации чисел Фибоначчи и отправки их в канал
func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch) // Закрываем канал
}

func main() {
	ch := make(chan int)

	go fibonacci(10, ch)

	// Чтение из канала
	for num := range ch {
		fmt.Printf("Число Фибоначчи: %d\n", num)
	}
}
