package exercise1

import (
	"fmt"
	"strings"
)

func Etl() {
	one_to_many := map[int][]string{
		1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  {"D", "G"},
		3:  {"B", "C", "M", "P"},
		4:  {"F", "H", "V", "W", "Y"},
		5:  {"K"},
		8:  {"J", "X"},
		10: {"Q", "Z"},
	}

	one_to_one := map[string]int{}
	for points, letters := range one_to_many {
		for _, letter := range letters {
			letter = strings.ToLower(letter)
			one_to_one[letter]=points
		}
	}

	for points, letter :=range one_to_one{
		fmt.Println(letter," : ",points)
	}
}
