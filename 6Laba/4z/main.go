package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

// Функция для увеличения счетчика
func increment(wg *sync.WaitGroup, useMutex bool) {
	defer wg.Done()

	for i := 0; i < 10000; i++ {
		if useMutex {
			mutex.Lock()
		}
		counter++
		if useMutex {
			mutex.Unlock()
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// Без мьютекса (Большая вероятность гонки данных)
	wg.Add(2)
	go increment(&wg, false)
	go increment(&wg, false)
	wg.Wait()
	fmt.Println("Счётчик без мьютекса:", counter)

	// С мьютексом (Без гонки данных)
	counter = 0
	wg.Add(2)
	go increment(&wg, true)
	go increment(&wg, true)
	wg.Wait()
	fmt.Println("Счётчик с мьютексом:", counter)
}
