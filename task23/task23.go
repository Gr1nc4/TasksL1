package main

import "fmt"

//Удалить i-ый элемент из слайса.

func main() {
	arr := []int{5, 6, 2, 12}
	fmt.Println(removeElem(arr, 3))
}

func removeElem(arr []int, i int) []int { //благодаря оператору среза в go достаточно просто удалить элемент из слайса
	arr = append(arr[:i], arr[i+1:]...)
	return arr
}
