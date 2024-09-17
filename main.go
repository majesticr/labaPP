package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Время: ", currentTime.String())
	var (
		a int     = 1
		b float64 = 3.14
		c string  = "word"
		d bool    = true
	)
	e := "Text"

	fmt.Println("int:", a)
	fmt.Println("float64:", b)
	fmt.Println("string:", c)
	fmt.Println("bool:", d)
	fmt.Println("short variable:", e)

}
