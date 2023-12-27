package main

import (
	"fmt"
	"sort"
	"strings"
)

// Функция для создания ключа анаграммы
func createAnagramKey(word string) string {
	word = strings.ToLower(word)
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// Функция поиска множеств анаграмм
func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	seen := make(map[string]bool)

	for _, word := range words {
		key := createAnagramKey(word)
		if _, exists := anagrams[key]; !exists {
			anagrams[key] = []string{}
		}
		wordLower := strings.ToLower(word)
		if !seen[wordLower] {
			anagrams[key] = append(anagrams[key], wordLower)
			seen[wordLower] = true
		}
	}

	// Удаление множеств с одним элементом и выбор ключа
	for key, group := range anagrams {
		if len(group) < 2 {
			delete(anagrams, key)
			continue
		}
		sort.Strings(group)
		firstWord := group[0]
		anagrams[firstWord] = group
		delete(anagrams, key)
	}

	return anagrams
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "один"}
	anagrams := findAnagrams(words)
	for key, group := range anagrams {
		fmt.Println("Ключ:", key, "Группа:", group)
	}
}
