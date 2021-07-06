package main

import (
	"fmt"
	"time"
)

func customSleep1(i int){
	<-time.After(time.Second * time.Duration(i))
}

func customSleep2(i int) {
	now := time.Now().Second()
	for time.Now().Second() - now < i{}
}

func main() {
	str := "Hello world"
	//customSleep1(3)
	customSleep2(3)
	fmt.Println(str)

}
