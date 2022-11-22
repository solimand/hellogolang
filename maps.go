package main

import (
	"fmt"
	"strings"
)

// This function returns a map of the words in the string s and the number of occurrences of the words in the string
func WordCount(s string) map[string]int {
	var words = strings.Fields(s)
	result := make(map[string]int)

	// FIND DUPLICATES
	for _, v := range words {
		_, ok := result[v]
		if ok {
			result[v] += 1
		} else {
			result[v] = 1
		}
	}

	fmt.Println(result)
	//return map[string]int{"x": 1}
	return result
}

// func main() {
// 	wc.Test(WordCount)
// }
