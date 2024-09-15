package main

import "fmt"

func main() {

	// создание переменных
	a := 50.5
	b := 12.4

	// вызов функций
	func1(a, b)
	func2(a, b)

}

// функция для сложения
func func1(x float64, y float64) {
	var z = x + y
	fmt.Println("x + y = ", z)
}

// функция для вычитания
func func2(x float64, y float64) {
	var z = x - y
	fmt.Println("x - y = ", z)
}
