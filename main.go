package main

import (
	"fmt"
	"time"
)

func main() {
	type MyDuration time.Duration
	d := MyDuration(100)

	fmt.Printf("%T", d)
}
