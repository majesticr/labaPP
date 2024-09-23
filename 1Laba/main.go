package main

import (
	"fmt"
	"time"
)

// 5.Вычисление разности и суммы чисел с плавающей точкой
func calculate(one, two float64) (sum float64, difference float64) {
	sum = one + two
	difference = one - two
	return
}

// 6.Вычисление среднего значения 3х чисел
func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func main() {

	// 1.Вывод текущего времени и даты
	currentTime := time.Now()
	fmt.Println("Время: ", currentTime.String())

	// 2.Создание переменных различных типов
	var (
		a int     = 1
		b float64 = 3.14
		c string  = "word"
		d bool    = true
	)

	// 3. Краткая форма объявления переменных
	e := "Text"
	f := true
	g := 44
	h := 44.44

	//Вывод переменных заданий 2 и 3
	fmt.Println("(a) int:", a)
	fmt.Println("(b) float64:", b)
	fmt.Println("(c) string:", c)
	fmt.Println("(d) bool:", d)
	fmt.Println("(e) short variable (e):", e)
	fmt.Println("(f) short variable (f):", f)
	fmt.Println("(g) short variable (g):", g)
	fmt.Println("(h) short variable (h):", h)

	// 4. Арифметические операции с двумя целыми числами
	one, two := 12, 21
	fmt.Println("Арифметические операции:")
	fmt.Printf("Сумма %d + %d = %d\n", one, two, one+two)
	fmt.Printf("Разность %d - %d = %d\n", one, two, one-two)
	fmt.Printf("Произведение %d * %d = %d\n", one, two, one*two)
	fmt.Printf("Деление (частное) %d / %d = %.2f\n", one, two, float64(one)/float64(two))

	// 5. Вывод вычисления разности и суммы чисел с плавающей точкой
	sum, difference := calculate(3.14, 6.3)
	fmt.Printf("Сумма: %.2f, Разность: %.2f\n", sum, difference)

	// 6. Вывод среднего значения 3х чисел
	avg := average(77.0, 88.0, 99.0)
	fmt.Printf("Ср.знач.: %.2f\n", avg)

}
