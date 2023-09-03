package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"hoge": 42,
		"fuga": 33,
	}

	age := m["John"]
	fmt.Println(age)

	age, ok := m["Jane"]
	fmt.Println(age, ok)
}

func even(n int) bool {
	return n%2 == 0
}
