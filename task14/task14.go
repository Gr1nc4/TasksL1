package main

import "fmt"

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int type ", v)
	case string:
		fmt.Println("string type", v)
	case bool:
		fmt.Println("bool type", v)
	case chan int:
		fmt.Println("channel type", v)
	default:
		fmt.Println("unknown type", v)
	}
}

func main() {
	do("124")
}
