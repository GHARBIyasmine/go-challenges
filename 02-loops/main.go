package main

import "fmt"

func main() {
	// Print the multiplication table from 1 to 10
	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println("")
	}

	// sum a hard-coded slice of 6 numbers and print the total
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice len: ", len(slice))
	fmt.Println("slice cap: ", cap(slice))
	var sum int = 0
	for _, value := range slice {
		sum += value
	}
	fmt.Println("sum: ", sum)

	// Using a condition-only loop, find the first power of 2 greater than 1000
	p2 := 1
	for p2 < 1000 {
		p2 *= 2
	}

	fmt.Println("First power of 2 > 1000: ", p2)
}
