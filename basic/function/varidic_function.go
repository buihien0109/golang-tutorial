package main

import "fmt"

func Hello(name string, number ...int) {
	fmt.Println("name : " + name)
	for index, value := range number {
		fmt.Printf("index = %d - value = %d\n", index, value)
	}
}

func Find(number int, numbers ...int) {
	fmt.Printf("type of nums is %T\n", numbers)
	found := false
	for index, value := range numbers {
		if number == value {
			fmt.Println(number, "found at index", index)
			found = true
		}
	}

	if !found {
		fmt.Println(number, "not found in ", numbers)
	}
	fmt.Println()
}
