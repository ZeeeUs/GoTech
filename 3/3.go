package main

import (
	"fmt"
)

func Squaring(arr* [5]int, ch chan int){
	for _, v := range arr{
		ch <- v * v
	}
}

func Sum(ch chan int) {
	t := 0
	for i := 0; i < 5; i++ {
		t += <-ch
	}
	fmt.Println(t)
}

func main(){
	intCh := make(chan int)
	arr := [5]int { 2, 4, 6, 8, 10 }
	go Squaring(&arr, intCh)
	Sum(intCh)
}