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
		after      = flag.Int("A", 0, "print +N lines after match")
		before     = flag.Int("B", 0, "print +N lines before match")
		context    = flag.Int("C", 0, "print ±N lines around match")
		count      = flag.Bool("c", false, "print count of matching lines")
		ignoreCase = flag.Bool("i", false, "ignore case")
		invert     = flag.Bool("v", false, "select non-matching lines")
		fixed      = flag.Bool("F", false, "fixed string match")
		lineNumber = flag.Bool("n", false, "print line number")
	)
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Usage: grep [options] <pattern> <file>")
		os.Exit(1)
	}
	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	matchCount := 0
	var beforeLines []string
	printAfter := 0

	for scanner.Scan() {
		line := scanner.Text()
		originalLine := line
		lineNum++

		if *ignoreCase {
			line = strings.ToLower(line)
			pattern = strings.ToLower(pattern)
		}

		match := strings.Contains(line, pattern)
		if *fixed {
			match = line == pattern
		}
		if *invert {
			match = !match
		}

		if match {
			matchCount++
			if *lineNumber {
				fmt.Printf("%d:", lineNum)
			}
			fmt.Println(originalLine)

			if *after > 0 || *context > 0 {
				printAfter = max(*after, *context)
			}
		} else {
			if printAfter > 0 {
				fmt.Println(originalLine)
				printAfter--
			}
		}

		if *before > 0 || *context > 0 {
			if len(beforeLines) >= max(*before, *context) {
				beforeLines = beforeLines[1:]
			}
			beforeLines = append(beforeLines, originalLine)
		}

		if match && (*before > 0 || *context > 0) {
			for _, bLine := range beforeLines[:len(beforeLines)-1] {
				fmt.Println(bLine)
			}
			beforeLines = beforeLines[:0]
		}
	}

	if *count {
		fmt.Println(matchCount)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
