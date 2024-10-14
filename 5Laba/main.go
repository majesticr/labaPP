package main

import (
	"fmt"
	"math"
)

// Person структура
type Person struct {
	name string
	age  int
}

// Метод для вывода информации о человеке
func (p Person) Info() {
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}

// Метод Birthday для увеличения возраста на 1 год
func (p *Person) Birthday() {
	p.age++
}

// Circle структура
type Circle struct {
	radius float64
}

// Метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Rectangle структура
type Rectangle struct {
	width, height float64
}

// Метод для вычисления площади прямоугольника
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Shape интерфейс с методом Area()
type Shape interface {
	Area() float64
}

// Функция, которая принимает срез интерфейсов Shape и выводит площадь каждого объекта
func PrintAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}

// Book структура для хранения информации о книге
type Book struct {
	title  string
	author string
}

// Stringer интерфейс с методом String()
type Stringer interface {
	String() string
}

// Реализация интерфейса Stringer для структуры Book
func (b Book) String() string {
	return fmt.Sprintf("Book Title: %s, Author: %s", b.title, b.author)
}

func main() {
	// Пример работы с Person
	p := Person{name: "Alexander", age: 20}
	p.Info()
	p.Birthday()
	p.Info()

	// Пример работы с Shape (Circle и Rectangle)
	c := Circle{radius: 6}
	r := Rectangle{width: 5, height: 4}

	shapes := []Shape{c, r}
	PrintAreas(shapes)

	// Пример работы с интерфейсом Stringer и структурой Book
	b := Book{title: "Harry Potter", author: "Joanne Rowling"}
	fmt.Println(b.String())
}
