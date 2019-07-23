package main

import "fmt"

func main() {
	fmt.Println("couting from 1  to 100")

	for i := 0; i <= 100; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

