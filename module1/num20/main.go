package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("тестовое сообщение"))
}

/*
Я использовал следующий подход:
Разделил исходную строку на слова.
Перевернул порядок слов в массиве.
Объединил слова обратно в строку.
*/
func reverseWords(str string) string {
	words := strings.Fields(str)
	//я думаю тут можно решить задачу эффективней при помощи обхода сразу с 2 сторон
	//но я решил для разнообразия сделать вспомогательный слайс
	reversedWords := make([]string, 0, len(words)) //сразу указываю капасити чтобы аллокатору проще жилось
	for i := len(words) - 1; i >= 0; i-- {
		reversedWords = append(reversedWords, words[i])
	}
	return strings.Join(reversedWords, " ")
}
