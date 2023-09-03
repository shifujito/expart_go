package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2, 3, 4, 5}

	dst := src[:0]
	fmt.Println(dst)
	fmt.Println(src)
	for _, v := range src {
		if even(v) {
			dst = append(dst, v)
		}
	}
	fmt.Println(dst)
	fmt.Println(src)
	for i := len(dst); i < len(src); i++ {
		src[i] = 0
	}
	fmt.Println(dst)
	fmt.Println(src)
}

func even(n int) bool {
	return n%2 == 0
}
