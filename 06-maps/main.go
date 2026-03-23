package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	phrase := "the quick brown fox jumps over the lazy dog the fox"
	words := strings.Fields(phrase)
	m := make(map[string]int)
	for _, word := range words {
		if v, ok := m[word]; ok {
			m[word] = v + 1
		} else {
			m[word] = 1
		}

	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, word := range keys {
		fmt.Printf("%s: %d\n", word, m[word])
	}

	var maxKeys []string
	maxValue := 0
	for k, v := range m {
		if v > maxValue {
			maxValue = v
			maxKeys = []string{k}
		} else if v == maxValue {
			maxKeys = append(maxKeys, k)
		}
	}

	fmt.Printf("Most frequent word(s): %v with count: %d\n", maxKeys, maxValue)

	delete(m, "the")
	fmt.Println(m)

}
