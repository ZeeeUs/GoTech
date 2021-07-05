package main

import "fmt"

func main() {
	a := []int{1, 4 , 8, 7, 10}
	b := []int{1, 8 , 4, 7, 55, 66, 10}
	cross := []int{}

	for _, v := range a{
		for _, val := range b{
			if v == val{
				cross = append(cross, val)
			}
		}
	}

	fmt.Printf("%v\n", cross)
}
