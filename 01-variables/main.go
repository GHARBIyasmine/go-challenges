package main

import "fmt"

func main() {
	const z float32 = -273.15

	var (
		c float32
		f float32
		k float32
	)
	c = 100
	f = (c * 9 / 5) + 32
	k = c + 273.15

	fmt.Printf("Temperature: %.0f°C", c)
	fmt.Printf("\nFahrenheit: %.2f°F", f)
	fmt.Printf("\nKelvin: %.2fK", k)
	fmt.Printf("\nAbsolute zero: %.2f°C", z)
}
