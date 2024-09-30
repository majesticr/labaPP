package main

import (
	"fmt"
	"myproject/mathutils"
	"myproject/stringutils"
)

func main() {
	// Факториал введенного числа
	var num int
	fmt.Print("Введите число для вычисления факториала: ")
	fmt.Scan(&num)
	fmt.Printf("Факториал числа %d = %d\n", num, mathutils.Factorial(num))

	// Переворот строки
	var str string
	fmt.Print("Введите строку для переворота: ")
	fmt.Scan(&str)
	fmt.Printf("Перевернутая строка: %s\n", stringutils.Reverse(str))

	// Создание массива и его вывод
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Массив:", arr)

	// Создание среза из массива
	slice := arr[:]
	fmt.Println("Срез из массива:", slice)

	// Добавление элемента в срез
	slice = append(slice, 6)
	fmt.Println("Срез после добавления элемента:", slice)

	// Удаление элемента из среза
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("Срез после удаления третьего элемента:", slice)

	// Создание среза строк и нахождение самой длинной строки
	strSlice := []string{"строка первая", "длинное слово", "элемент кода", "управление системой", "проверка кода"}
	longest := strSlice[0]
	for _, s := range strSlice {
		if len(s) > len(longest) {
			longest = s
		}
	}
	fmt.Printf("Самая длинная строка: %s\n", longest)
}
