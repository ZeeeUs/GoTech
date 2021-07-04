package main

import (
	"fmt"
)

func main() {
	arr := []int{1, -55, 66, 88, -99, 3, 14, -1, 0, 33}
	fmt.Printf("%+v\n", arr)
	//sort.Search(arr) // Сортировка встроенным методом языка, используется quickSort
	fmt.Printf("%+v\n", arr)
}
