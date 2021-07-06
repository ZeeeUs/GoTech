package main

import "fmt"

func reverse(str string) string {
	arr := []rune(str)
	buf := []rune{}
	for i := len(str) - 1; i >= 0; i-- {
		buf = append(buf, arr[i])
	}

	return string(buf)
}

func main() {
	fmt.Println(reverse("hello"))
}
