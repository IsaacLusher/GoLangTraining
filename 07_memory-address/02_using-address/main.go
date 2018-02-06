package main

import "fmt"

const metresToYards float64 = 1.09361

func main(){
	var metres float64
	fmt.Print("Enter meters ran: ")
	fmt.Scan(&metres)
	yards := metres * metresToYards
	fmt.Println(metres, " metres is ", yards, " yards.")
}
