package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для генерации случайных чисел
func generateNumbers(ch chan int) {
	for i := 0; i < 5; i++ { // Ограничиваем количество чисел до 5
		ch <- rand.Intn(10)
		time.Sleep(2 * time.Second)
	}
	close(ch) // Закрываем канал после завершения генерации чисел
}

// Функция для определения четности/нечетности числа
func checkEvenOdd(ch chan int, msgChan chan string) {
	for num := range ch {
		if num%2 == 0 {
			msgChan <- fmt.Sprintf("%d - четное", num)
		} else {
			msgChan <- fmt.Sprintf("%d - нечетное", num)
		}
	}
	close(msgChan) // Закрываем канал после завершения обработки всех чисел
}

func main() {
	numChan := make(chan int)
	msgChan := make(chan string)

	go generateNumbers(numChan)
	go checkEvenOdd(numChan, msgChan)

	// Чтение сообщений о четности/нечетности
	for msg := range msgChan {
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}
