package main

import (
	"fmt"
	"sync"
)

type someStruct struct {
	cnt int
	mut sync.Mutex
}

func count(i int, ch *chan bool, str *someStruct){
	str.mut.Lock()
	str.cnt = 0
	for t := 1; t <= 5; t++{
		str.cnt++
		fmt.Printf("Gorutine: %v - %v\n", i, str.cnt)
	}
	str.mut.Unlock()
	*ch <- true
}

func main() {
	ch := make(chan bool)
	for i:= 1; i < 5; i++ {
		go count(i, &ch, &someStruct{})
	}

	for i:= 1; i < 5; i++ {
		<-ch
	}

	fmt.Println("The end")
}
