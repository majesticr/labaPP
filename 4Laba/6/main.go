package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Введите количество элементов в массиве:")
	fmt.Scan(&n)

	// Создание массива
	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")

	// Считывание элементов массива
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// Вывод массива в обратном порядке
	fmt.Println("Array in reverse order:")
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}
