package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

//Смысл работы этого паттерна в том, что если у вас есть класс и его интерфейс не совместим с кодом вашей системы,
//то что бы разрешить этот конфликт, мы не изменяем код этого класса, а пишем для него адаптер.
//Другими словами Adapter адаптирует существующий код к требуемому интерфейсу (является переходником).

// К чему требуется адаптировать
type Animal interface {
	Speak()
}

type People struct{}

func (p People) Speak() {
	fmt.Println("Hello!")
}

//Адаптируемый
type Dog interface {
	Bark()
}

type Сorgi struct{}

func (c Сorgi) Bark() {
	fmt.Println("Wow!")
}

//Адаптер
type AdapterDog struct {
	dog Сorgi
}

func (a AdapterDog) Speak() {
	a.dog.Bark()
}

func main() {
	var dog Animal
	dog.Speak()
}
