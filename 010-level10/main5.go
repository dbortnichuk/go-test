package main

import (
	"fmt"
)

func main5() {
	q := make(chan int)
	c := gen5(q)

	receive5(c, q)

	close(q)

	fmt.Println("about to exit")
}

func gen5(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			c <- i
		}

		close(c)
		q <- 0
		
	}()

	
	return c
}

func receive5(c, q <-chan int) {

	for {
		select {
		case v := <-c:
			fmt.Println("c:", v)
		case v := <-q:
			fmt.Println("q:", v)
			return
		}
	}
}
