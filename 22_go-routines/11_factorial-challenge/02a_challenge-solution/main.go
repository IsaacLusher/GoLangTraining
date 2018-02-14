package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println("Total:", <-accum(f))
}

func factorial(n int) chan int {
	c := make(chan int)
	go func() {
		for i := n; i > 0; i-- {
			c <- i
		}
		close(c)
	}()
	return c
}

func accum(c chan int) chan int {
	out := make(chan int)
	go func() {
		a := 1
		for x := range c {
			a *= x
		}
		out <- a
		close(out)
	}()
	return out
}
