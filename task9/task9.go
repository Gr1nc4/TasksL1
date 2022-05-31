package main

import "fmt"

/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout. */

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr := []int{1, 2, 3, 4, 5}

	go in(arr, ch1)
	go out(ch2, ch1)

	for val := range ch2 {
		fmt.Println(val)
	}
}

//Пишем значения в первый канал
func in(arr []int, ch1 chan int) {
	for _, val := range arr {
		ch1 <- val
	}
	close(ch1)
}

//пишем значения*2 из первого канала во второй
func out(ch2 chan int, ch1 chan int) {
	for val := range ch1 {
		ch2 <- val * 2
	}
	close(ch2)
}
