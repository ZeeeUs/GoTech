package main

import "fmt"

func reverseString(str string) []string {
	strLocal := []rune(str)
	var buf []string
	var buf2 []string
	anchor := 0
	for i, v := range strLocal {
		if v == ' ' {
			buf = append(buf, string(strLocal[anchor:i]))
			anchor = i + 1
		}
	}
	buf = append(buf, string(strLocal[anchor:]))

	for i := len(buf) - 1; i >= 0; i-- {
		buf2 = append(buf2, buf[i])
	}

	return buf2
}

func main(){
	fmt.Println(reverseString("snow dog sun"))
}
