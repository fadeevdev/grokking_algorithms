package main

import "fmt"

func main() {
	arr := []int{6, 12, 323, 3123, 4124, 56322, 566546}
	fmt.Println(binary_search(arr, 12))
	fmt.Println(binary_search(arr, 29))
	fmt.Println(binary_search(arr, 566546))
}

func binary_search(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
