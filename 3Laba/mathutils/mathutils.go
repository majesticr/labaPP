package mathutils

// Создание функции Factorial для нахождения факториала числа
func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}
