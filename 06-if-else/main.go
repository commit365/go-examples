package main

import "fmt"

func main() {
	fmt.Println("If/Else control flow")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	fmt.Println("If statement with no else clause")
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	fmt.Println("If statement with OR || operator")
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 8 or 7 is even")
	}

	fmt.Println("If statement with local variable declaration")
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println("\nNotes")
	fmt.Println("  You don't need parentheses around conditionals.")
	fmt.Println("  There is not ternary operator like `conditional ? true : false`")
}
