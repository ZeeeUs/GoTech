package main

import (
	"fmt"
	"time"
)

func Worker(ch chan interface{}, num int){
	for i := range ch {
		fmt.Printf("Worker #%v: %v\n", num, i)
	}
}

func main() {
	countOfWorker := 4
	ch := make(chan interface{})

	for i := 0; i < countOfWorker; i++ {
		time.Sleep(time.Second)
		go Worker(ch, i+1)
	}

	for t := 0; t < countOfWorker; t++{
		ch <- "hello world"
		ch <- 2*2
		ch <- "Moscow " + "Russia"
		ch <- time.Now()
	}

	close(ch)
}
