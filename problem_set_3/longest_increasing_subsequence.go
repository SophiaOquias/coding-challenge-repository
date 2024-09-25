package main 

import "fmt"

func lengthOfLIS(num []int) int {
	length := len(num)
	lastIndex := length - 1
	longestLength := 0

	for i := 0; i < length; i++ {
		current := num[i]
		sum := 0
		for j := i; j < length; j++ {
			if current <= num[j] {
				sum++
				current = num[j]
			}

			if j == lastIndex && longestLength < sum{
				longestLength = sum
			}
		}
	}

	return longestLength
}

func main() {
	inputNumbers := []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := lengthOfLIS(inputNumbers)
	fmt.Println(result) // Output: 4
}