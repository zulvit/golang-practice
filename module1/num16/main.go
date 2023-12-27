package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3, 1, 3, 514, 134, 1, 3, 513, 4, 51, 3, 31, 678}
	sort.Ints(slice)
	fmt.Println(slice)
}
