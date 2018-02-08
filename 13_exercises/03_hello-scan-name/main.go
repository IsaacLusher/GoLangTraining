package main

import "fmt"

func main() {
	var name1 string
	var name2 string
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&name1)
	fmt.Print("Please enter your surname: ")
	fmt.Scan(&name2)
	fmt.Println("Hello", name1, name2)
}
