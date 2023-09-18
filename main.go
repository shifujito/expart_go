package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	bob := person{}
	fmt.Println(bob)
	// fmt.Println(bob == nil)
}
