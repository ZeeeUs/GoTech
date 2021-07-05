package main

import "fmt"

func set(arr []string) []string {
	buf := make(map[string][]int)
	set := []string{}

	for i, v := range arr {
		buf[v] = append(buf[v], i)
	}

	for key, _ := range buf {
		set = append(set, key)
	}

	return set
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	Set := set(words)

	fmt.Printf("All words: %v\n", words)
	fmt.Printf("Set: %v\n", Set)
}
