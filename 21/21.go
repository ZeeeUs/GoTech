package main

import "fmt"

func main(){
	arr := []int{2, 4, 8, 16, 32, 64}
	ch := make(chan int)

	go func(chan int, []int){
		for i:= 0; i < len(arr); i++ {
			ch <- arr[i]
		}
		close(ch)
	}(ch, arr)

	for v := range ch{
		fmt.Println(v)
	}
}
