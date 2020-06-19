package main

import (
	"fmt"
)

// func produceTo1(c chan<- int){
// 	for i:= 1; i <= 3; i++ {
// 		c <- i
// 	}
// 	close(c)
// }

func source1() <-chan int {
	sc := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	go func() {
		for i := 1; i <= 3; i++ {
			sc <- i
		}
		//wg.Done()
		close(sc)
	}()
	//wg.Wait()
	
	return sc
}

func source2() <-chan int {
	sc := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	go func() {
		for i := 4; i <= 5; i++ {
			sc <- i
		}
		close(sc)
		//wg.Done()
	}()
	//wg.Wait()
	
	return sc
}

func merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for {
			out <- <-in1
		}
	}()

	go func() {
		for {
			out <- <-in2
		}
	}()
	return out
}

func sink(source <-chan int) {

	var v int
	var ok bool

	for {
		v, ok = <- source
		if(ok){
			fmt.Println(v)
		} else {
			return
		}
	}
}

func main7() {
	s1 :=  source1()
	s2 :=  source2()

	out := merge(s1, s2)

	sink(out)

	fmt.Println("about to exit")
	// s2 :=  source2()
	// m := merge(s1, s2)
	// sink(m)
}
