package main

import "fmt"

func main() {
	arr := []int{11, 22, 33, 44, 55}
	low := 0
	high := len(arr) - 1
	target := 566

	fmt.Println(binSearchRec(arr, low, high, target))
}

func binSearchRec(arr []int, low, high, target int) int {

	if low > high {
		fmt.Println("Target doesnt exists")
		return -1
	}

	mid := (low + high) / 2

	if target == arr[mid] {
		return mid
	} else if target < arr[mid] {
		return binSearchRec(arr, low, mid-1, target)
	} else {
		return binSearchRec(arr, mid+1, high, target)
	}

	return mid

}
