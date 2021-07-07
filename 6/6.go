package main

import (
	"fmt"
	"math/rand"
	"time"
)

func someFunc() {
	for {
		i := rand.Intn(100)
		if i == 55 {
			break
		}
	}
	fmt.Println("True")
}

func main() {
	go someFunc()
	time.Sleep(time.Second * 2)
}
