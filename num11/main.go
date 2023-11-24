package main

import (
	"fmt"
)

// intersect принимает два множества (map[тип_ключа]bool) и возвращает их пересечение.
func intersect(set1, set2 map[int]bool) map[int]bool {
	intersection := make(map[int]bool)
	for key := range set1 {
		if set2[key] {
			intersection[key] = true
		}
	}
	return intersection
}

func main() {
	set1 := map[int]bool{1: true, 2: true, 3: true}
	set2 := map[int]bool{2: true, 3: true, 4: true}

	result := intersect(set1, set2)
	fmt.Println("Пересечение множеств:", result)
}
