package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		fields    = flag.String("f", "", "fields to select")
		delimiter = flag.String("d", "\t", "delimiter to use")
		separated = flag.Bool("s", false, "only lines with delimiter")
	)
	flag.Parse()

	fieldIndexes, err := parseFields(*fields)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing fields:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		columns := strings.Split(line, *delimiter)
		var selectedColumns []string
		for _, index := range fieldIndexes {
			if index < len(columns) {
				selectedColumns = append(selectedColumns, columns[index])
			}
		}
		fmt.Println(strings.Join(selectedColumns, *delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
		os.Exit(1)
	}
}

func parseFields(fields string) ([]int, error) {
	var indexes []int
	for _, field := range strings.Split(fields, ",") {
		var index int
		_, err := fmt.Sscanf(field, "%d", &index)
		if err != nil {
			return nil, err
		}
		indexes = append(indexes, index-1) // Adjust for zero-based indexing
	}
	return indexes, nil
}
