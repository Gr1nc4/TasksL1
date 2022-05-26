package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int) //Создаем канал, в который будем писать наши данные
	sum := 0             //Переменная для подсчета суммы

	for _, val := range arr {
		go mult1(ch, val) //конкурентно пишем данные в канал
	}
	for i := 0; i < len(arr); i++ {
		sum += <-ch //Считываем данные из канала и прибавляем к имеющейсы у нас сумме
	}
	fmt.Println(sum)
}

//Записываем результат умножение в канал
func mult1(ch chan int, a int) {
	ch <- a * a
}
