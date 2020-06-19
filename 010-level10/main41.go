package main

import (
	"fmt"
)

func main41() {
	c := gen41()
	receive41(c)

	fmt.Println("about to exit")
}

func receive41(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func gen41() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 4; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
