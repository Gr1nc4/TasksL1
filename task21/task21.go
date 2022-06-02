package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

//Смысл работы этого паттерна в том, что если у вас есть класс и его интерфейс не совместим с кодом вашей системы,
//то что бы разрешить этот конфликт, мы не изменяем код этого класса, а пишем для него адаптер.
//Другими словами Adapter адаптирует существующий код к требуемому интерфейсу (является переходником).

//какой-то внешний сервис, который нужно адаптировать
type AnalyticalDataService interface {
	SendXmlData()
}
type Xml struct {
}

func (doc Xml) SendXmlData() {
	fmt.Println("Отправка XML")
}

//Наша система, которая работает с json
type Json struct {
}

func (doc Json) ConvertToXml() string {
	return "XML"
}

//Наш адаптер, который позволяет реализовывать методы внешнего сервиса
type JsonDocAdapter struct {
	json *Json
}

func (adapter JsonDocAdapter) SendXmlData() {
	adapter.json.ConvertToXml()
}

func main() {
	var json JsonDocAdapter
	json.SendXmlData()
}
