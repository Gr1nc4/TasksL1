package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func testUnique(s string) bool {
	s = strings.ToLower(s)
	m := make(map[string]int)
	for _, value := range s { //проходимся по нашей строке и провряем имеется ли такой символ(ключ) в мапе
		if _, ok := m[string(value)]; ok {
			return false //если встретили такой ключ - значит символы в строке не уникальны
		}
		m[string(value)]++
	}
	return true
}

func main() {
	str := "abcd"
	str1 := "abCdefAaf"
	str2 := "aabcd"
	fmt.Println(str, "-", testUnique(str))
	fmt.Println(str1, "-", testUnique(str1))
	fmt.Println(str2, "-", testUnique(str2))
}
