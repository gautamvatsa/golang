package main

import "fmt"

func main() {
	odd := [5]int{3, 5, 7, 11, 13}

	fmt.Println("no are ",odd)
	for v :=range odd{
	if v/2==0 {
	fmt.Println("No is even")
	} else {
	fmt.Println("No is odd")
}
                 }
	 }
