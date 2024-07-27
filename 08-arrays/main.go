package main

import "fmt"

func main() {
	fmt.Println("Empty 5 int array")
	var a [5]int
	fmt.Println("emp:", a)

	fmt.Println("")
	fmt.Println("Set a value at an index using the array[index] = value syntax")
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("")
	fmt.Println("`len` returns the length of an array")
	fmt.Println("len:", len(a))

	fmt.Println("")
	fmt.Println("Declare and initialize an array in one line")
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	fmt.Println("")
	fmt.Println("Compiler counts the number of elements for you with ...")
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	fmt.Println("")
	fmt.Println("Specify the index with : and the elements in between will be zeroed")
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	fmt.Println("")
	fmt.Println("Compose types to build multi-dimensional data structures")
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	fmt.Println("")
	fmt.Println("Create and initialize multi-dimensional arrays at once")
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
