# Problem Set 3: Longest Increasing Subsequence  

## Problem Description
Given an unsorted array of integers, determine the length of the longest increasing subsequence. 

## Solution Overview
The solution starts by first initializing a `longestLength` variable to `0`; this represents the length of the longest increasing subsequence. We then iterate through the array of integers in a nested loop. Each iteration starts at different indices of the array. The starting index represents the index of the number at the beginning of a subsequence. As we iterate through the array, the subsequent numbers are checked if they are lesser than or equal to the most recent number in the subsequence. If it is lesser than or equal to the number in the subsequence a `length` variable is incremented. At the end of every subsequence, if `longestLength` is lesser than `length`, then `length` replaces `longestLength` in value. Finally, once all possible starting indices are checked, the function then returns `longestLength`.

## Instructions to Run the Code
### Java Instructions 
1. `cd` into the `problem_set_2` folder
2. Compile the program by running: 
```
javac LongestIncreasingSubsequence.java
```
3. Run the program with: 
```
java LongestIncreasingSubsequence
```
### Go Instructions
1. `cd` into the `problem_set_2` folder 
2. Run by using: 
```
go run .
```