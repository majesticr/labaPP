package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для расчета факториала
func factorial(n int) {
	time.Sleep(1 * time.Second) // Задержка
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Printf("Факториал %d: %d\n", n, result)
}

// Функция для генерации случайных чисел
func generateRandomNumbers(count int) {
	time.Sleep(2 * time.Second) // Задержка
	for i := 0; i < count; i++ {
		fmt.Printf("Случайное число: %d\n", rand.Intn(10))
	}
}

// Функция для вычисления суммы числового ряда
func sumSeries(n int) {
	time.Sleep(3 * time.Second) // Задержка
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("Сумма числового ряда до %d: %d\n", n, sum)
}

func main() {
	go factorial(3)
	go generateRandomNumbers(3)
	go sumSeries(3)

	time.Sleep(5 * time.Second)
}
