package main 

import "fmt"

type Node struct {
	item rune 
	next *Node
}

type Stack struct {
	top *Node
}

func (stack *Stack) push(bracket rune) {
	node := &Node{item: bracket, next: stack.top}

	stack.top = node
}

func (stack *Stack) pop() {
	stack.top = stack.top.next
}

func (stack Stack) print() {
	current := stack.top

	for current != nil {
		fmt.Printf("%c\n", current.item)
		current = current.next
	}

}

func (stack Stack) isEmpty() bool {
	if(stack.top == nil) {
		return true
	}

	return false
}

func isOpenBracket(bracket rune) bool {
	if bracket == '(' || bracket == '[' || bracket == '{' {
		return true
	}

	return false 
}

func getMatch(bracket rune) rune {
	switch bracket {
		case ')': return '(' 
		case ']': return '['
		case '}': return '{'
		default: return 0 
	}
}

func isValid(s string) bool {
	stack := &Stack{}

	for i := 0; i < len(s); i++ {
		current := rune(s[i])

		if isOpenBracket(current) {
			stack.push(current)
		} else {
			if stack.isEmpty() {
				return false 
			} 

			if getMatch(current) != stack.top.item {
				return false
			}

			stack.pop()
		}
	}

	return true
}

func main() {
	inputString := "([)]"
	result := isValid(inputString)
	fmt.Println(result) 
}