package main

import "fmt"

func wordCount(words []string) map[string]int {
	
	counts := make(map[string]int)

	
	for _, word := range words {

		counts[word]++
	}

	return counts
}

func main() {

	inputWords := []string{"go", "is", "go", "fast"}

	result := wordCount(inputWords)

	fmt.Println("Word Counts:", result)
}