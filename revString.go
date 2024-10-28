package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

func reverseString(str string) string {
	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024)

	for {
		fmt.Println("Enter at least 3 words in a sentence: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		wordCount := countWords(text)

		if wordCount < 3 {
			fmt.Println("Not enough words! Try again.")
			continue
		}

		fmt.Println("Number of Words:", wordCount)
		fmt.Println("Entered Text:", text)
		fmt.Println("Reversed Text:", reverseString(text))
		break
	}
}
