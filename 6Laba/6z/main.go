package main

import (
	"fmt"
	"strings"
	"sync"
)

// Воркер, который реверсирует строку
func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Закрытие воркера, когда он завершает выполнение
	for job := range jobs {
		fmt.Printf("Воркер %d обрабатывает строку: %s\n", id, job)
		// Реверс строки и отправка результата в канал
		reversed := reverseString(job)
		results <- fmt.Sprintf("Воркер %d: %s -> %s", id, job, reversed)
	}
}

// Функция для реверса строки
func reverseString(s string) string {
	return strings.Join(reverse(strings.Split(s, "")), "")
}

// Вспомогательная функция для реверса массива строк
func reverse(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	var wg sync.WaitGroup

	jobs := make(chan string, 10)    // Канал для задач
	results := make(chan string, 10) // Канал для результатов

	// Количество воркеров
	numWorkers := 3

	// Запуск пула воркеров
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Массив строк для обработки
	stringsToProcess := []string{"hello", "world", "workers", "number", "table"}

	// Отправка задач в канал jobs
	go func() {
		for _, str := range stringsToProcess {
			jobs <- str
		}
		close(jobs) // Закрытие канала jobs после отправки всех задач
	}()

	// Чтение результатов
	go func() {
		wg.Wait()      // Ожидание завершения всех воркеров
		close(results) // Закрытие канала результатов
	}()

	// Вывод результатов
	for result := range results {
		fmt.Println(result)
	}
}
