package main

import "fmt"

func greatest(n ...int) int {
	var biggest int
	for _, v := range n {
		if v > biggest {
			biggest = v
		}
	}
	return biggest
}

func main() {
	biggestNumber := greatest(3847, 454, 32, 46859, 648, 26, 879, 6451, 547, 976, 1200)
	fmt.Println(biggestNumber)
}
