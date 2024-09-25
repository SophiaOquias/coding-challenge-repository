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

func main() {
	stack := &Stack{}

	stack.push('A')
	stack.push('B')
	stack.push('C')

	stack.print()
}