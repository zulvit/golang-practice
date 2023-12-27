package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 2)
	slice[0] = "a"
	slice[1] = "a"

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
