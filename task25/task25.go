package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func Sleep(i int) {
	<-time.After(time.Duration(i) * time.Second)
}

func main() {
	i := 5
	fmt.Println("Я")
	Sleep(i)
	fmt.Println("Поспал")
}
