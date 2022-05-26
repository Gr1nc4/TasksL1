package main

import "fmt"

func main() {
	a := 4
	b := 6

	a, b = b, a

	fmt.Println("a =", a, "b =", b)
}
