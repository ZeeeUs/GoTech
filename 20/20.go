package main

import "fmt"

func main() {
	slice := []string{"a", "a"}
	func(slice []string) { // анонимня функция
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Printf("%v\n", slice)
	}(slice) // передаём не по ссылке, а по значению, т.е. слайс копируется и мы работаем с локальным слайсом
	// Т.к. передовали по значению, то исходный слайс не изменился
	fmt.Printf("%v\n", slice)
}
