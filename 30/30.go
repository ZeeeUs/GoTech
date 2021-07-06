package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//i := 3

	arr = arr[:len(arr)-1]
	fmt.Println(arr)
}
