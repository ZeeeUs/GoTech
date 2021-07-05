package main

import "fmt"

func TypeOfStruct(val interface{}) {
	switch val.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Int")
	case bool:
		fmt.Println("Bool")
	case struct{}:
		fmt.Println("Struct")
	case chan int:
		fmt.Println("Channel")
	default:
		fmt.Println("This type is undefined")
	}
}

func main(){
	TypeOfStruct("Hello")
	TypeOfStruct(123)
	TypeOfStruct(true)
	TypeOfStruct(make(chan int))
}
