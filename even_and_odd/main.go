package main

import "fmt"

func main() {
	mySlice := []int16{0,1,2,3,4,5,6,7,8,9,10}

	for _, number := range mySlice {
		if  (number % 2 == 0) {
			fmt.Printf("%v is even \n", number)
		} else {
			fmt.Printf("%v is odd \n", number)
		}
	}
}