package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Функция для сравнения строк в зависимости от флагов
func compareStrings(a, b string, numeric bool, column int) bool {
	if column > 0 {
		aColumns := strings.Fields(a)
		bColumns := strings.Fields(b)
		if len(aColumns) >= column && len(bColumns) >= column {
			a, b = aColumns[column-1], bColumns[column-1]
		}
	}

	if numeric {
		aNum, errA := strconv.ParseFloat(a, 64)
		bNum, errB := strconv.ParseFloat(b, 64)
		if errA == nil && errB == nil {
			return aNum < bNum
		}
	}
	return a < b
}

func main() {
	column := flag.Int("k", 0, "column to sort by")
	numeric := flag.Bool("n", false, "sort by numeric value")
	reverse := flag.Bool("r", false, "reverse the result")
	unique := flag.Bool("u", false, "unique lines only")

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: mysort [options] <file>")
		os.Exit(1)
	}
	filename := flag.Arg(0)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Сортировка
	sort.SliceStable(lines, func(i, j int) bool {
		return compareStrings(lines[i], lines[j], *numeric, *column)
	})

	if *reverse {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}

	if *unique {
		uniqueLines := make([]string, 0, len(lines))
		prevLine := ""
		for _, line := range lines {
			if line != prevLine {
				uniqueLines = append(uniqueLines, line)
				prevLine = line
			}
		}
		lines = uniqueLines
	}

	// Вывод результатов
	for _, line := range lines {
		fmt.Println(line)
	}
}
