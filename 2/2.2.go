package main

import (
	"fmt"
)

func Calc (arr [5]int, ch chan int) {
	for _, v := range arr {
		ch <- v * v
	}
	close(ch)
}

func main()  {
	arr := [5]int { 2, 4, 6, 8, 10 }
	ch := make(chan int)
	go Calc(arr, ch)

	for i := range ch {
		fmt.Printf("%v \n", i)
	}
}
