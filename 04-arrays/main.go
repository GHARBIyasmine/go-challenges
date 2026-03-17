package main

import "fmt"

func main() {
	arr := [7]int{2000, 2500, 3000, 3500, 4000, 4500, 5000}
	total := 0
	highest := 0
	highestIndex := 0
	lowest := arr[0]
	lowestIndex := 0
	for i, v := range arr {
		fmt.Printf("Day %v: %v\n", i+1, v)
		total += v
		if v > highest {
			highest = v
			highestIndex = i
		}
		if v < lowest {
			lowest = v
			lowestIndex = i
		}

	}

	fmt.Println("Total steps: ", total)
	fmt.Println("Average steps: ", total/len(arr))
	fmt.Printf("Best day: Day %v (%v steps)\n", highestIndex+1, highest)
	fmt.Printf("Worst day: Day %v (%v steps)\n", lowestIndex+1, lowest)
}
