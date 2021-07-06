package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 4

	fmt.Println("Исходный слайс: ", arr)

	arr = append(arr[:i-1], arr[i:]...)
	fmt.Println(arr)

	arr[i] = arr[len(arr)-1]
	arr[len(arr)-1] = 0
	arr = arr[:len(arr)-1]

	fmt.Println(arr)
}
