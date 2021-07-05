package main

import "fmt"

func readFromA(a []int, ch chan<- int) {
	for _, val := range a{
		ch <- val
	}
	close(ch)
}

func Equal(b []int, ch <-chan int) []int {
	cross := []int{}
	for val := range ch {
		for _, v := range b {
			if val == v {
				cross = append(cross, val)
			}
		}
	}
	return cross
}

func main(){
	ch := make(chan int)
	a := []int{1, 4 , 8, 7, 10}
	b := []int{1, 8 , 4, 7}

	go readFromA(b, ch)
	fmt.Printf("First slice: %v\n", a)
	fmt.Printf("Second slice: %v\n", b)
	fmt.Printf("Cross: %v\n", Equal(a, ch))
}
