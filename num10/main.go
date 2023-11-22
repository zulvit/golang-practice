package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := groupTemperatures(temperatures, 10)

	for k, v := range groups {
		fmt.Printf("Группа %v: %v\n", k, v)
	}
}

// groupTemperatures группирует значения, основываясь на шаге в 10 градусов и возвращает мапу
func groupTemperatures(temps []float64, step float64) map[float64][]float64 {
	groups := make(map[float64][]float64)

	for _, temp := range temps {
		key := math.Round(temp/step) * step
		groups[key] = append(groups[key], temp)
	}
	return groups
}
