package main

import "fmt"

// Для того чтобы в go работать с символами строки, нужно приводить ее к слайсу рун т.к.
// изначально строка в go это массив байт. А из-за того что разные символы могут занимать разное количество байт
// итерация по строке из кириллицы и латиныцы с будут давать разный результат.

func createHugeString(size int) (s string) {
	for i := 0; i < size; i++ {
		s += "Б"
	}
	return
}

func someFunc() {
	v := createHugeString(1 << 10)
	//создаем слайс рун
	a := []rune(v)
	//берем слайс до 10 элемента
	justString := v[:10]
	//Привеоди наш слайс к string и так же берем срез до 10 элемента
	justStringRune := string(a[:10])

	fmt.Println(justString)     //В этом случае мы выведем срез до 10 байт, но в
	fmt.Println(justStringRune) //В этом случае выводим срезд до 10 символов
}

func main() {
	someFunc()
}
