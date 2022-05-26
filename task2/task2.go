package task2

//func main() {
//	arr := []int{2, 4, 6, 8, 10}
//	ch := make(chan int) //Создаем канал, в который будем писать наши данные
//
//	for _, val := range arr {
//		go mult(ch, val) //конкурентно пишем данные в канал
//	}
//	for i := 0; i < len(arr); i++ {
//		fmt.Println(<-ch) //выводим то что лежит в канале
//	}
//}
//
////Записываем результат умножение в канал
//func mult(ch chan int, a int) {
//	ch <- a * a
//}
