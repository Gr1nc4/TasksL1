package task1

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от
//родительской структуры Human (аналог наследования).

type Human struct {
	Name string
	Age  uint
}
type Action struct {
	Id    string
	Human // Встраиваение Human в Action. Action наследует поля Human
}

//func main() {
//	NewAction := Action{
//		Id: "askdyiu12y3",
//		Human: Human{
//			Name: "John",
//			Age:  32,
//		},
//	}
//	fmt.Println(NewAction)
//}
