package main

import "fmt"

// Функция для вычисления среднего возраста
func averageAge(people map[string]int) float64 {
	totalAge := 0
	for _, age := range people {
		totalAge += age
	}
	return float64(totalAge) / float64(len(people))
}

func main() {
	// Пример карты
	people := map[string]int{
		"Alexandr": 10,
		"Evgeniy":  20,
		"Roman":    30,
		"Pavel":    40,
	}

	// Вывод среднего возраста
	fmt.Printf("Average Age: %.2f\n", averageAge(people))
}
