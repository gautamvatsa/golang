package main

import "fmt"

func main() {
	t := 2
	switch {
	case t == 1 :
		fmt.Println("T = 1 ")
	case t == 2 :
		fmt.Println("T = 2")
	default:
		fmt.Println("T = nothing")
	}
}

