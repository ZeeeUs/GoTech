package main

import (
	"fmt"
	"time"
)

func Write(arr *[10]string, ch chan string){
	for _, val := range arr {
		ch <-val
	}
}

func Read(a int, ch chan string){
	for i := 0; i < a; i++{
		fmt.Println(<-ch)
	}
}

func main() {
	arr := [10] string {"One", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	channel := make(chan string)
	go Write(&arr, channel)
	go Read(len(arr), channel)

	time.Sleep(time.Second)
}
