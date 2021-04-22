package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fmt.Println(s)
	words := strings.Fields(s)
	result := make(map[string]int)
	var i int
	for i < len(words) {
		counter := 0
		var j int
		for j < len(words) {
			if words[i] == words[j] {
				counter++
				result[words[j]] = counter
			}
			j++
		}
		i++
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
