package main

import "fmt"

func main() {
	arr := make([]int, 100)
	fmt.Printf("%v\n", arr)
	fmt.Printf("Capacity: %v\n", cap(arr))
	fmt.Printf("Length: %v\n", len(arr))
}
