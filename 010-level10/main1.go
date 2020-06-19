package main

import (
	"fmt"
)

func main1() {
	c := make(chan int)

	go func () {
		c <- 42
	}()
	

	fmt.Println(<- c)
}

func main2() {
	c := make(chan int, 2)

	c <- 42
	c <- 43
	
	fmt.Println(<- c)
	fmt.Println(<- c)
}