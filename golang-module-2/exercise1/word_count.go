package exercise1

import (
	"fmt"
	"regexp"
	"strings"
)

func Word_coun() {
	
	text := "\"That's the password: 'PASSWORD 123'!\", cried the Special Agent.\nSo I fled."
	re := regexp.MustCompile(`[a-zA-Z0-9]+(?:'[a-zA-Z0-9]+)?`)
	words := re.FindAllString(text, -1)
	// fmt.Println(words)
	frequency := map[string]int{}

	for _, word := range words {
		word = strings.ToLower(word)
		frequency[word]++
	}

	for word, count := range frequency{
		fmt.Println(word," : ", count)
	}
}
