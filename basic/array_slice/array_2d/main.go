package main

import "fmt"

func main() {
	var a [2][3]int
	fmt.Println(a)

	a[0][0] = 1
	a[0][1] = 2
	a[0][2] = 3
	a[1][0] = 4
	a[1][1] = 5
	a[1][2] = 6
	fmt.Println(a)

	// In array
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%d ", v2)
		}
		fmt.Printf("\n")
	}
}
