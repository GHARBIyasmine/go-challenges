package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	phrase := "racecar 🚀 Go"
	fmt.Println("len in bytes ", len(phrase))
	fmt.Println("num of runes ", utf8.RuneCountInString(phrase))
	fmt.Println("num of words ", len(strings.Fields(phrase)))

	rune_slice := []rune(phrase)
	fmt.Println(rune_slice)

	for i, j := 0, len(rune_slice)-1; i < j; i, j = i+1, j-1 {
		rune_slice[i], rune_slice[j] = rune_slice[j], rune_slice[i]
	}
	fmt.Println(rune_slice)

	r := []rune("racecar")
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println(string(r) == "racecar")

	fmt.Println(strings.Replace(phrase, "Go", "Golang", 1))

}
