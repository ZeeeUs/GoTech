package main

import "fmt"

func check(s string) bool {
	str := []rune(s)
	buf := map[rune]int{}

	for _, v := range str{
		buf[v] += 1
	}

	for _, v := range buf{
		if v != 1 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(check("Hello"))
}
