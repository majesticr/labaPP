package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Считывание строки
	fmt.Println("Введите строку:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Вывод строки в верхнем регистре
	fmt.Println(strings.ToUpper(input))
}
