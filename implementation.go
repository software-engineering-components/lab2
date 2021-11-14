package lab2

import (
	"fmt"
)

type StackNode struct {
	data string
	next * StackNode
}
func getStackNode(data string, top * StackNode) * StackNode {
	// return new StackNode
	return &StackNode {
		data,
		top,
	}
}
type MyStack struct {
	top * StackNode
	count int
}
func getMyStack() * MyStack {
	// return new MyStack
	return &MyStack {
		nil,
		0,
	}
}
// Returns the number of element in stack
func(this MyStack) size() int {
	return this.count
}
func(this MyStack) isEmpty() bool {
	if this.size() > 0 {
		return false
	} else {
		return true
	}
}
// Add a new element in stack
func(this *MyStack) push(data string) {
	// Make a new stack node
	// And set as top
	this.top = getStackNode(data, this.top)
	// Increase node value
	this.count++
}
// Add a top element in stack
func(this *MyStack) pop() string {
	var temp string = ""
	if this.isEmpty() == false {
		// Get remove top value
		temp = this.top.data
		this.top = this.top.next
		// Reduce size
		this.count--
	}
	return temp
}
// Used to get top element of stack
func(this MyStack) peek() string {
	if !this.isEmpty() {
		return this.top.data
	} else {
		return ""
	}
}

// Check operator
func isOperator(text byte) bool {
	if text == '+' || text == '-' ||
	   text == '*' || text == '/' ||
	   text == '^' || text == '%' {
		return true
	}
	return false
}
// Check operands
func isOperands(text byte) bool {
	if (text >= '0' && text <= '9') ||
	   (text >= 'a' && text <= 'z') ||
	   (text >= 'A' && text <= 'Z') {
		return true
	}
	return false
}
// Converting the given postfix expression to
// infix expression
func postfixToInfix(postfix string) (string, error) {
	// Get the size
	var size int = len(postfix)
	// Create stack object
	var s * MyStack = getMyStack()
	// Some useful variables which is using
	// of to storing current result
	var auxiliary string = ""
	var op1 string = ""
	var op2 string = ""
	var isValid bool = true
	for i := 0 ; i < size && isValid ; i++ {
		// Check whether given postfix location
		// at [i] is an operator or not
		if isOperator(postfix[i]) {
			// When operator exist
			// Check that two operands exist or not
			if s.size() > 1 {
				op1 = s.pop()
				op2 = s.pop()
				auxiliary = "(" + op2 + string(postfix[i]) + op1 + ")"
				s.push(auxiliary)
			} else {
				isValid = false
			}
		} else if isOperands(postfix[i]) {
			// When get valid operands
			auxiliary = string(postfix[i])
			s.push(auxiliary)
		} else {
			// Invalid operands or operator
			isValid = false
		}
	}
	if isValid == false {
		// When have something wrong
		return postfix, fmt.Errorf("Invalid postfix : ", postfix)
	} else {
		// Display calculated result
		fmt.Println(" Postfix : ", postfix)
		// fmt.Println(" Infix   : ", s.pop())
		return s.pop(), nil
	}
}
