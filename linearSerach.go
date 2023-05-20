package main

import "fmt"

func main() {
	arr := []int{44, 33, 22, 99, 17}
	target := 99
	fmt.Println(lSearch(arr, target))
}

func lSearch(arr []int, target int) int {
	for key, val := range arr {
		if val == target {
			return key
		}
	}
	return -1
}
