// package main
// import ("fmt")

// func main() {
//   fmt.Println("Hello all!")
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, b, c, d int = 1, 3, 5, 7

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var (
// 		a int
// 		b int    = 1
// 		c string = "hello"
// 	)

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)

// }

package main

import "fmt"

func main() {

	// var cars = [3]string{"Vovlo", "Hundai", "Verna"}
	// cars1 := [...]string{"Rajat", "Mohadikar"}
	// fmt.Println(cars1)
	// cars1[1] = "Honda"
	// fmt.Println(cars1)

	// fmt.Println(cars)

	// arr1 := [...]string{1: "Ram", 2: "Raj"}
	// fmt.Println("arr1=", arr1)
	// fmt.Println(len(arr1))
	// fmt.Println(cap(arr1))

	// arr2 := [6]int{88, 33, 9, 10, 77, 45}

	// myslice1 := arr2[1:6]
	// fmt.Println(len(myslice1))
	// fmt.Println(cap(myslice1))

	// //mySlice2 := []int{33, 22, 44, 11, 55}

	// mySlice3 := make([]int, 5, 55)
	// fmt.Printf("capacity=%d\n", cap(mySlice3))
	// fmt.Printf("length=%d\n", len(mySlice3))
	// fmt.Println("myslice3=%v", mySlice3)

	// day := 44

	// switch day {
	// case 1, 66:
	// 	fmt.Println("Monday")

	// case 2, 55:
	// 	fmt.Println("Tueday")

	// case 3, 44:
	// 	fmt.Println("Wed")

	// case 4:
	// 	fmt.Println("Thursday")

	// default:
	// 	fmt.Print("Default Tag")
	// }

	// arr1 := [5]string{"Apple", "Mango", "Banana"}
	// arr2 := [5]string{"Lenovo", "Asus", "HP"}

	// for i := 1; 1 < len(arr1); i++ {
	// 	for j := 0; j < len(arr2); j++ {
	// 		fmt.Println(arr1[i], arr2[i])
	// 	}
	// }

	// fruits := [3]string{"Mango", "Apple", "Banana"}
	// for idx, val := range fruits {
	// 	//fmt.Println("Index=%d\tValue=%d", idx, val)
	// 	fmt.Printf("%v\t%v\n", idx, val)
	// }

	// fruits := [3]string{"apple", "orange", "banana"}
	// for idx, val := range fruits {
	// 	fmt.Printf("%v\t%v\n", idx, val)
	// }

	// var a = map[string]string{"brand": "Ford", "Model": "Mustang", "Year": "1964"}
	// b := map[string]int{"students": 10, "City": 2, "Country": 1}
	// fmt.Printf("a\t%v\n", a)
	// fmt.Printf("b\t%v", b)

	// var a = map[string]string{"Brand": "Ford", "Year": "1964", "Type": "Petrol"}

	// val1, ok1 := a["brand"]
	// val2, ok2 := a["color"]
	// _, ok3 := a["Year"]
	// val4, ok4 := a[" "]

	// fmt.Println(ok1, val1)
	// fmt.Println(ok2, val2)
	// fmt.Println(ok3)
	// fmt.Println(ok4, val4)

	// var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	// val1, ok1 := a["brand"] // Checking for existing key and its value
	// val2, ok2 := a["color"] // Checking for non-existing key and its value
	// val3, ok3 := a["day"]   // Checking for existing key and its value
	// _, ok4 := a["model"]    // Only checking for existing key and not its value

	// fmt.Println(val1, ok1)
	// fmt.Println(val2, ok2)
	// fmt.Println(val3, ok3)
	// fmt.Println(ok4)

	a := map[string]int{"One": 1, "Two": 2, "Three": 3}

	for key, val := range a {
		fmt.Printf("Key= %s, Value=%d", key, val)
	}

  var b = []string             // defining the order
  b = append(b, "one", "two", "three", "four")

  for _, element := range b{
    fmt.Printf("%v  %v",element, a[element])
  }

}
