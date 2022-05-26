package main

import "fmt"

func main() {
	arr := []int{5, 6, 2, 12}
	fmt.Println(removeElem(arr, 3))
}

func removeElem(arr []int, i int) []int {
	arr = append(arr[:i], arr[i+1:]...)
	return arr
}
