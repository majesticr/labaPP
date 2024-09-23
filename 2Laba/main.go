package main

import (
	"fmt"
	"strings"
)

// 1. Проверка четности/нечетности числа
func checkEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}

// 2. Проверка знака числа
func checkSign(num int) string {
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

// 3. Вывод чисел от 1 до 10 с помощью цикла for
func Numbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 4. Функция подсчета длины строки
func stringLength(s string) int {
	return len(s)
}

// 5. Структура Rectangle и вычисление площади прямоугольника
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 6. Функция для вычисления среднего значения двух чисел
func averagecount(a, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	// 1. Четность/нечетность числа
	var num int
	fmt.Print("Введите число для проверки четности/нечетности: ")
	fmt.Scan(&num)
	checkEvenOdd(num)

	// 2. Проверка знака числа
	fmt.Print("Введите число для проверки знака: ")
	fmt.Scan(&num)
	fmt.Println("Знак:", checkSign(num))

	// 3. Вывод чисел от 1 до 10
	fmt.Println("Числа от 1 до 10:")
	Numbers()

	// 4. Определение длины строки
	var input string
	fmt.Print("Введите строку для вычисления ее длины: ")
	fmt.Scan(&input)
	fmt.Println("Длина строки:", stringLength(strings.TrimSpace(input)))

	// 5. Площадь прямоугольника
	var width, height float64
	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scan(&width)
	fmt.Print("Введите высоту прямоугольника: ")
	fmt.Scan(&height)
	rect := Rectangle{Width: width, Height: height}
	fmt.Println("Площадь прямоугольника:", rect.Area())

	// 6. Ср.знач 2х чисел
	var num1, num2 int
	fmt.Print("Введите 2 числа для вычисления среднего: ")
	fmt.Scan(&num1, &num2)
	fmt.Println("Среднее значение:", averagecount(num1, num2))
}
