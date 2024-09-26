# Problem Set 2: Valid Parentheses  

## Problem Description
Check if a string of brackets (i.e. '(', ')', '[', ']', '{', '}') are closed by its matching bracket and is closed in the correct order.  

## Solution Overview
The solution starts by iterating over the string of brackets. The bracket is checked if it is an opening bracket (i.e. '(', '[', '{'). If it is an opening bracket, it is pushed to a stack. Otherwise,it is a closing bracket. If it is a closing bracket, the stack is first checked if it is empty. If it is empty, that means that there is no matching opening bracket, which means an invalid string of brackets; the function then returns false. If the stack is not empty, then the stack is checked if the top of the stack contains the matching opening bracket of the closed bracket. For example, if the closed bracket is ')', the top of the stack is checked if it is a '('. If the top of the stack does not contain the matching bracket, this means an invalid string of brackets and the function then returns false. Otherwise, the stack is popped and we move on to the next bracket in the string. 

## Instructions to Run the Code
### Java Instructions 
1. `cd` into the `problem_set_1` folder
2. Compile the program by running: 
```
javac ValidParentheses.java
```
3. Run the program with: 
```
java ValidParentheses
```
### Go Instructions
1. `cd` into the `problem_set_1` folder 
2. Run by using: 
```
go run .
```