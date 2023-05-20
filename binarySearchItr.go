package main

import "fmt"

func main() {
	arr := []int{12, 14, 33, 43, 74}

	target := 889
	fmt.Println(binSearchItr(arr, target))
}

func binSearchItr(arr []int, target int) int {

	low := 0
	high := len(arr) - 1
	var mid int
	for low < high {
		mid = (low + high) / 2

		if target == arr[mid] {
			return mid
		} else if target < arr[mid] {
			high--
		} else {
			low++
		}
	}

	fmt.Println("Target doesnt exists")
	return -1

}
