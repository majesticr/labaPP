package main

import "fmt"

func main() {
	// Карта с именами людей и их возрастами
	people := map[string]int{
		"Alexandr": 20,
		"Oleg":     30,
		"Maksim":   40,
	}

	// Добавление нового человека
	people["Vladimir"] = 50

	// Вывод
	for name, age := range people {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
}
