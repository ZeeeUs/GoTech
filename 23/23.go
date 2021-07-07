package main

import "fmt"

func binarySearch(arr []int, i int) int {
	start := 0
	end := len(arr) - 1
	mid := len(arr) / 2

	for start <= end {
		val := arr[mid]

		if val == i {
			return mid
		}

		if val > i {
			end = mid - 1
			mid = (start + end) / 2
			continue
		}

		start = mid + 1
		mid = (start + end) / 2
	}

	return 0
}

func main() {
	arr := []int{1, 2, 6, 13, 15, 25, 33, 40, 41, 42, 50, 55}
	i := 6
	fmt.Println(binarySearch(arr, i))
}
