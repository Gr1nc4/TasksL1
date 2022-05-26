package main

import "fmt"

/* Реализовать пересечение двух неупорядоченных множеств. */

func main() {
	arr1 := []int{2, 5, -9, 21, 14, 3, 4}
	arr2 := []int{-2, 5, 8, 7, 3}
	newArr := make([]int, 0)
	//Сравниваем каждый элемент первого массива с каждым элементов второго массива
	//Если находим совпадение кладем в newArr
	for _, val1 := range arr1 {
		for _, val2 := range arr2 {
			if val1 == val2 {
				newArr = append(newArr, val1)
			}
		}
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(newArr)
}
