package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a

}

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c Counter) Value() int {
	return c.count
}

func (c *Counter) Reset() {
	c.count = 0
}

func main() {
	a, b := 5, 10
	fmt.Println("Before swap: a =", a, "b =", b)
	swap(&a, &b)
	fmt.Println("After swap: a =", a, "b =", b)

	counter := Counter{}
	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Println("Counter value:", counter.Value())
	counter.Reset()
	fmt.Println("Counter value after reset:", counter.Value())
}
