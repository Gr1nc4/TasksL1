package main

import "fmt"

//имеется последовательность строк - (cat, cat, dog, cat, tree)
//создать для нее собственное множество.

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	var res []string
	newArr := make(map[string]int)
	//т.к. ключи в мапе уникальны, будем складывать в мапу значения слайса в качестве ключей
	for _, val := range arr {
		if newArr[val] != 1 {
			newArr[val] = 1
			res = append(res, val)
		}
	}
	fmt.Println(res, newArr)
}
