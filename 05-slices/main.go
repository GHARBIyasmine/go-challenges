package main

import "fmt"

func main() {
	var s []string
	s = append(s, "Buy groceries")
	s = append(s, "Write tests")
	s = append(s, "Read a book")
	s = append(s, "Fix bug #42")
	fmt.Println("cap", cap(s))
	fmt.Println("len", len(s))
	for i, v := range s {
		fmt.Printf("%d: %s\n", i, v)
	}
	fmt.Println("Removing item: ", s[2])
	s = append(s[0:2], s[3:]...)
	fmt.Println("cap", cap(s))
	fmt.Println("len", len(s))
	for i, v := range s {
		fmt.Printf("%d: %s\n", i, v)
	}
	ss := make([]string, len(s))
	copy(ss, s)
	fmt.Println("cap", cap(ss))
	fmt.Println("len", len(ss))
	ss[0] = "Buy groceries and cook dinner"
	if ss[0] != s[0] {
		fmt.Println("Original unchanged: true")
	} else {
		fmt.Println("Original unchanged: false")
	}
}
