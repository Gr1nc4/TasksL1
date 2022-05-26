package main

/* Разработать программу, которая будет последовательно
отправлять значения в канал,а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться. */

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("Нужно указать время работы в секундах")
	}
	n, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan int)
	go write(ch)
	go read(ch)

	time.Sleep(time.Duration(n) * time.Second)
	close(ch)

}

func read(ch chan int) {
	for {
		fmt.Println("из канала пришло", <-ch)
	}

}
func write(ch chan int) {
	i := 0
	for {
		ch <- i
		fmt.Println("Записали в канал", i)
		time.Sleep(500 * time.Millisecond)
		i++
	}
}
