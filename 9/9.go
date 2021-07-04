package main

import "fmt"

func Iterator(nums ...int) <-chan int {
	firstCh := make(chan int)
	go func() {
		defer close(firstCh)
		for _, val := range nums {
			firstCh <- val
		}
	}()

	return firstCh
}

func Squaring(firstCh <-chan int) <-chan int {
	secondCh := make(chan int)
	go func() {
		defer close(secondCh)
		for val := range firstCh {
			secondCh <- val * val
		}
	}()

	return secondCh
}

func main() {
	first := Iterator(2, 4, 8, 16)
	second := Squaring(first)

	for val := range second {
		fmt.Printf("%d\n", val)
	}
}
