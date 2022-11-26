package main

import "fmt"

func main() {
	number := 50
	guess := 300

	if guess < number {
		fmt.Println("Too low")
	}

	if guess > number {
		fmt.Println("Too high")
	}

	if guess == number {
		fmt.Println("Equal")
	}

	if guess > 1 && guess < 100 {
		fmt.Println("Yes in range")
	} else {
		fmt.Println("Not in range")
	}

	returnTrue()

	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("nor one or two")
	}

	i := 10

	switch {
	case i <= 10:
		fmt.Println("less than or equal")
		fallthrough
	case i <= 20:
		fmt.Println("less than 20")
	default:
		fmt.Println("greater than 20")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
