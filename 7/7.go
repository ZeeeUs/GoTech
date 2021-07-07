package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mymap = map[int]int{}
	m := &sync.Mutex{}

	for i := 0; i < 5; i++{
		go func (mymap map[int]int, t int, m *sync.Mutex) {
			for j := 0; j < 5; j++ {
				m.Lock()
				mymap[t*10+j]++
				m.Unlock()
			}
		}(mymap, i, m)
	}
	time.Sleep(time.Millisecond * 10)
	m.Lock()
	fmt.Println("counters result", mymap)
	m.Unlock()
}
