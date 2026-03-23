package main

import "fmt"

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	defer fmt.Println("--- calculation complete ---")
	fmt.Println("Addition:", add(5, 3))
	fmt.Println("Subtraction:", subtract(5, 3))
	fmt.Println("Multiplication:", multiply(5, 3))

	result, err := divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result)
	}

	fmt.Println("Sum of 1, 2, 3, 4:", sum(1, 2, 3, 4))
}
