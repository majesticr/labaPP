package main

import "fmt"

func main() {
	// Карта людей с возрастами
	people := map[string]int{
		"Alexandr": 20,
		"Oleg":     30,
		"Maksim":   40,
	}

	// Имя для удаления
	nameToDelete := "Oleg"

	// Удаление записи по имени
	delete(people, nameToDelete)

	// Вывод оставшихся людей из карты
	for name, age := range people {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
}
