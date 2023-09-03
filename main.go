package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Name string

	m     sync.RWMutex
	count int
}

func (c *Counter) Increment() int {
	c.m.Lock()
	defer c.m.Unlock()
	c.count++
	return c.count
}

func (c *Counter) View() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.count
}

func main() {
	c := &Counter{
		Name: "Access",
	}
	fmt.Println(c.Increment())
	fmt.Println(c.View())
}
