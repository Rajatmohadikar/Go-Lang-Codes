package main

import "fmt"

func main() {

	// *****
	// ****
	// ***
	// **
	// *
	// 	size := 5
	// 	for i := size; i > 0; i-- {
	// 		for j := 0; j < i; j++ {
	// 			fmt.Print("*")
	// 		}
	// 		fmt.Println()
	// 	}

	// ******
	//  *****
	//   ****
	//    ***
	//     **
	//      *
	size := 5
	for i := size; i > 0; i-- {
		for k := i; k < size; k++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
