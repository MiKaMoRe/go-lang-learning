package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	wordsMap := make(map[string]int)

	for _, word := range words {
		value, ok := wordsMap[word]

		if ok {
			wordsMap[word] = value + 1
		} else {
			wordsMap[word] = 1
		}
	}
	return wordsMap
}

func main() {
	fmt.Println(WordCount("hello world beautiful world"))
}
