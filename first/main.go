package main

import "fmt"

func addWordCount(wordCount map[string]int, word string) {
	if count, ok := wordCount[word]; ok {
		wordCount[word] = count + 1
	} else {
		wordCount[word] = 1
	}
}

func countUniqueChars(s string) int {
	seen := make(map[byte]bool)
	count := 0

	for i := 0; i < len(s); i++ {
		if !seen[s[i]] {
			count++
			seen[s[i]] = true
		}
	}

	return count
}

func countUniqueSubstring(s string) int {
	ans := 0
	emptyMap := make(map[string]int)
	if s == "" {
		return 0
	}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substring := s[i:j]
			addWordCount(emptyMap, substring)
		}
	}
	for word, _ := range emptyMap {
		ans += countUniqueChars(word)
	}
	return ans
}
func main() {
	s := "ABC"
	fmt.Println(countUniqueSubstring(s))
	s = "ABA"
	fmt.Println(countUniqueSubstring(s))
}
