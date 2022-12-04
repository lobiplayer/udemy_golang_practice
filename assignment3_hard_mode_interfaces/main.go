package main

import (
	"fmt"
	"io"
	"os"
)

// func main() {
// 	fmt.Println(os.Args[1])
// }

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {

		io.Copy(os.Stdout, f)
	}
}
