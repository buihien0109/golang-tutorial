package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// BREAK
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	// CONTINUE
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Nested for loops
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// LABEL
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}

	}
}
