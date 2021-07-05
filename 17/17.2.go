package main

import (
	"fmt"
	"reflect"
)

func typeOfStruct(v interface{}) {
	// подробнее https://golang.org/pkg/reflect/
	valType := reflect.TypeOf(v)
	fmt.Println(valType)
}

func main() {
	typeOfStruct("Hello")
	typeOfStruct(123)
	typeOfStruct(true)
	typeOfStruct(make(chan int))
}