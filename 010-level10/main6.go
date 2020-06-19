package main

import (
	"fmt"
)

func main6() {
	c := make(chan int)

	go func(){

		c <- 1
		// for i := 1; i <= 4; i++{
		// 	c <- i
		// }
	}()

	v, ok := <- c
	fmt.Println(v, ok)

	close(c)
	
	v, ok = <- c
	fmt.Println(v, ok)

	v, ok = <- c
	fmt.Println(v, ok)
}