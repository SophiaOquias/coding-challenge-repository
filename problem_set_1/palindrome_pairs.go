package main

import "fmt"

func reverseString(s string) string {
	var reverse = ""
	var start = len(s) - 1

	for i := start; i >= 0; i-- {
		reverse = reverse + string(s[i])
	}

	return reverse; 
}

func isPalindrome(s string) bool {
	var reverse = reverseString(s) 

	if s == reverse {
		return true
	}

	return false
}

func palindromePairs(words []string) [][]int {
	var pairs [][]int 

	for i := 0; i < len(words); i++ {
		for j:= 0; j < len(words); j++ {
			if i != j {
				var concatinated = words[i] + words[j]
				if isPalindrome(concatinated) {
					var pair = []int{i, j}
					pairs = append(pairs, pair) 
				}
			}
		}
	}

	return pairs; 
}

func main() {
	inputWords := []string{"bat", "tab", "cat"}
	result := palindromePairs(inputWords)
	fmt.Println(result) // Output: [[0 1] [1 0]]
}