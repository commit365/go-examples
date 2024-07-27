package main

import "fmt"

func main() {
	fmt.Println("A single condition for loop")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("A classic initial/condition/after for loop")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("A range for loop")
	for i := range 3 {
		fmt.Println("range", i)
	}

	fmt.Println("A for loop with break")

	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("A for loop with continue")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
