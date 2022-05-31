package main

//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
//которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("Нужно указать колличество воркеров")
	}
	n, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan int)
	defer close(ch)
	//запуск воркеров в горутинах
	for i := 1; i <= n; i++ {
		go workerDo(i, ch)
	}
	//Сигналы ос пишем в канал, в нашем случае мы ждем SIGINT
	osSig := make(chan os.Signal, 1)
	defer close(osSig)
	signal.Notify(osSig, syscall.SIGINT)

	func() {
		var i int
		for {
			i++
			select {
			//Если получили данные из osSig выходим из функции
			case <-osSig:
				return
			default:
				time.Sleep(time.Millisecond * 500)
				ch <- i
			}
		}
	}()
}
func workerDo(id int, ch chan int) {
	for {
		fmt.Printf("Воркер #%d делает работу %d \n", id, <-ch)
	}
}
