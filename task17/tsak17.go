package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка

func main() {
	arr := []int{-6, -2, 0, 1, 8, 22} //берем отсортированный массив
	if search(arr, -2) {
		fmt.Println("Элемент найден")
	} else {
		fmt.Println("Элемент отсутствует")
	}

}

func search(arr []int, target int) bool {
	var min int
	max := len(arr) - 1
	//разбиваем наш массив пополам и проверяем наличие искомого элемента в каждой части
	for min <= max {
		mid := (max + min) / 2

		switch {
		case arr[mid] == target:
			return true
		case arr[mid] > target:
			max = mid - 1
		case arr[mid] < target:
			min = mid + 1
		}
	}
	return false
}
