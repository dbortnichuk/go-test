package main

import (
	"fmt"
)

func main4() {
	c := make(chan int)
	//receive4(chan<- int(c))

	go gen(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func gen(c chan<- int) {

	for i := 0; i < 4; i++ {
		c <- i
	}
	close(c)
}

func receive4(chan<- int){

}
