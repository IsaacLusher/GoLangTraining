package main

import (
	"fmt"
)

func main() {
	r := repeat()
	c := factorial(r)
	for n := range c {
		fmt.Println(n)
	}
}

func repeat() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 1; j < 11; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()
	return out
}
