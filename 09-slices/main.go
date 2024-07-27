package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("")
	fmt.Println("Uninitialized slice equals to nil and has length 0.")
	var s []string
	fmt.Println("unint:", s)
	fmt.Println("s == nil", s == nil)
	fmt.Println("len(s) == 0", len(s) == 0)

	fmt.Println("")
	fmt.Println("Use make to build an empty slice with non-zero length")
	s = make([]string, 3)
	fmt.Println("emp:", s)
	fmt.Println("len", len(s))
	fmt.Println("cap", cap(s))

	fmt.Println("")
	fmt.Println("Set and get slices just like arrays")
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("")
	fmt.Println("Get length with len")
	fmt.Println("len:", len(s))

	fmt.Println("")
	fmt.Println("Use append to append one or more values")
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	fmt.Println("")
	fmt.Println("Make slice with same length and copy values to it")
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	fmt.Println("")
	fmt.Println("Slice the slice with slice[low:high] syntax (excludes high)")
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	fmt.Println("")
	fmt.Println("Declare and initialize a variable for slice in a single line")
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	fmt.Println("")
	fmt.Println("Use a utility function for slices")
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	fmt.Println("")
	fmt.Println("Create multi-dimensional data structures with slices")
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
