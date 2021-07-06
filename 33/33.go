package main

import (
	"fmt"
	"time"
)

func parity(ch1, ch2 chan int) {
	num := <-ch1
	if num%2 == 0 {
		ch2 <- num
	}
}

func show(ch2 chan int) {
	fmt.Println(<-ch2)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr := []int{12, 15, 6, 7, 33, 69, 80, 2}

	for i := 0; i < len(arr); i++ {
		go parity(ch1, ch2)
		go show(ch2)
		ch1 <- arr[i]
	}

	time.Sleep(time.Millisecond)
}
