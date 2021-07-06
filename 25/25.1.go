package main

import (
	"fmt"
	"sync"
	"time"
)

type someOtherStruct struct {
	cnt int
	mut sync.Mutex
}

func (c *someOtherStruct) Count (){
	c.mut.Lock()
	c.cnt++
	c.mut.Unlock()
}

func (c *someOtherStruct) PrintResult() {
	fmt.Println(c.cnt)
}

func main() {
	counter := someOtherStruct{
		cnt: 0,
		mut: sync.Mutex{},
	}

	for i:=0; i < 5; i++{
		go func() {
			counter.Count()
			counter.PrintResult()
		}()
	}

	time.Sleep(time.Millisecond * 10)
}
