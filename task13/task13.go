package main

import "fmt"

//Поменять местами два числа без создания временной переменной.
func main() {
	a := 4
	b := 6
	a, b = b, a
	fmt.Println("a =", a, "b =", b)
}
