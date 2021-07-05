package main

import "fmt"

func main(){
	arr := []float32{-20, 0, 5, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -32.5}
	steps := make(map[int][]float32)

	for _, val := range arr {
		index := int(val/10)*10
		steps[index] = append(steps[index], val)
	}

	fmt.Printf("%v\n", steps)

}
