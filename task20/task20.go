package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	splitstr := strings.Split(str, " ")       //разбиваем нашу строчку по пробелами кладем слова в массив
	arr := make([]string, 0)                  //создаем массив куда будем складывать наши слова в обратном порядке
	for i := len(splitstr) - 1; i >= 0; i-- { //проходимся по нашему заспличенному массиву с конца и кладем элементы в arr
		arr = append(arr, splitstr[i])

	}
	fmt.Println("Исходная строка", str)
	fmt.Println("Развернутая строка ", strings.Join(arr, " "))
}
