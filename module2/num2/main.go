package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(s string) (string, error) {
	runes := []rune(s)
	var result []rune
	escaped := false

	for i, val := range runes {
		if val == '\\' && !escaped {
			escaped = true
			continue
		}

		if unicode.IsDigit(val) && !escaped {
			if i == 0 || unicode.IsDigit(runes[i-1]) {
				return "", fmt.Errorf("некорректная строка")
			}
			currentNum, err := strconv.Atoi(string(val))
			if err != nil {
				return "", err
			}
			repeat := strings.Repeat(string(runes[i-1]), currentNum-1)
			result = append(result, []rune(repeat)...)
		} else {
			result = append(result, val)
			escaped = false
		}
	}

	return string(result), nil
}

func main() {
	testStrings := []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`}

	for _, str := range testStrings {
		unpacked, err := unpackString(str)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Println(unpacked)
		}
	}
}
