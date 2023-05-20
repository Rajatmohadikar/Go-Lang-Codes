package main

import "fmt"

func main() {
	arr := []int{44, 33, 1, 11, 24}

	fmt.Println(bSort(arr))
}

func bSort(arr []int) []int {

	n := len(arr)

	if n == 0 || n == 1 {
		return arr
	} else {

		for i := 0; i < n-1; i++ {
			flag := 0
			for j := 0; j < n-i-1; j++ {
				if arr[j] > arr[j+1] {
					temp1 := arr[j+1]
					arr[j] = arr[j+1]
					arr[j+1] = temp1

					flag = 1
				}
			}
			if flag == 0 {
				break
			}
		}

	}

	return arr
}
