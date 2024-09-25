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

// func isPalindrome(s string) bool {

// }

// func palindromePairs(words []string) [][]int {

// }

func main() {
	// inputWords := []string{"bat", "tab", "cat"}
	// result := palindromePairs(inputWords)
	// fmt.Println(result) // Output: [[0 1] [1 0]]

	reverse := reverseString("hello")

	fmt.Println(reverse)
}