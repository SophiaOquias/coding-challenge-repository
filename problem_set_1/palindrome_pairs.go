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

// func palindromePairs(words []string) [][]int {

// }

func main() {
	// inputWords := []string{"bat", "tab", "cat"}
	// result := palindromePairs(inputWords)
	// fmt.Println(result) // Output: [[0 1] [1 0]]

	if isPalindrome("racecar") {
		fmt.Println("Palindrome!")
	} else {
		fmt.Println("Not Palindrome!")
	}
}