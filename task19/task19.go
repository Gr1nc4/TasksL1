package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	str := "главрыба!"
	fmt.Println(reverse(str))
}

func reverse(str string) string {
	result := ""
	sl := []rune(str)                   //Приводим нашу строку к слайсу рун
	for i := len(sl) - 1; i >= 0; i-- { //с конца слайса проходимся и кладем элементы в result
		result += string(sl[i])
	}
	return result
}
