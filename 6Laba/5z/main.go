package main

import (
	"fmt"
)

// Структура для запросов к калькулятору
type CalcRequest struct {
	A, B   int
	Op     string
	Result chan int
}

func calculator(reqChan chan CalcRequest) {
	for req := range reqChan {
		switch req.Op {
		case "+":
			req.Result <- req.A + req.B
		case "-":
			req.Result <- req.A - req.B
		case "*":
			req.Result <- req.A * req.B
		case "/":
			req.Result <- req.A / req.B
		}
	}
}

func main() {
	reqChan := make(chan CalcRequest)

	go calculator(reqChan)

	// Запросы
	for _, req := range []CalcRequest{
		{A: 2, B: 3, Op: "+", Result: make(chan int)},
		{A: 7, B: 2, Op: "-", Result: make(chan int)},
		{A: 10, B: 2, Op: "*", Result: make(chan int)},
		{A: 40, B: 2, Op: "/", Result: make(chan int)},
	} {
		reqChan <- req
		fmt.Printf("Результат операции %d %s %d = %d\n", req.A, req.Op, req.B, <-req.Result)
	}
}
