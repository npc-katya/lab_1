package main

import "fmt"

func main() {

	// создание переменных
	a := 7
	b := 3

	// арифметические операции
	c := a + b
	d := a - b
	e := a * b
	f := a / b
	g := a % b

	// вывод результатов
	fmt.Printf("a + b:  %d + %d = %d\n", a, b, c)
	fmt.Printf("a - b:  %d - %d = %d\n", a, b, d)
	fmt.Printf("a * b:  %d * %d = %d\n", a, b, e)
	fmt.Printf("a * b:  %d / %d = %d\n", a, b, f)
	fmt.Println("a % b: ", a, "%", b, "=", g)

}
