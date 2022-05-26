package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{3, 4, 32, -5, 20}
	qsort(arr)
	fmt.Println(arr)

}

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	l, r := 0, len(a)-1

	dot := rand.Int() % len(a)

	a[dot], a[r] = a[r], a[dot]

	for i, _ := range a {
		if a[i] < a[r] {
			a[l], a[i] = a[i], a[l]
			l++
		}
	}
	a[l], a[r] = a[r], a[l]
	qsort(a[:l])
	qsort(a[l+1:])

	return a
}
