package main

import "fmt"

func main() {
	str := "главрыба!"
	fmt.Println(reverse(str))
}

func reverse(str string) string {
	result := ""
	sl := []rune(str)
	for i := len(sl) - 1; i >= 0; i-- {
		result += string(sl[i])
	}
	return result
}
