package main

import (
	"fmt"
	"time"
)

func calc (arr* [5]int) {
	for i, v := range arr {
		arr[i] = v * v
	}
}

func main()  {
	arr := [5]int { 2, 4, 6, 8, 10 }
	go calc(&arr)
	time.Sleep(time.Second)

	fmt.Printf("%v \n", arr)
}
