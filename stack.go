package stack

import (
	"sync"
)

// thread safe stack

type Stack struct {
	top    *node
	length int
	mu     sync.Mutex
	Cond   *sync.Cond
}

type node struct {
	value interface{}
	prev  *node
}

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0, sync.Mutex{}, sync.NewCond(&sync.Mutex{})}
}

// Number of items in stack
func (this *Stack) Len() int {
	this.mu.Lock()
	defer this.mu.Unlock()
	return this.length
}

// Return the first item in the stack but dont remove it
func (this *Stack) Peek() interface{} {
	this.mu.Lock()
	defer this.mu.Unlock()

	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Return the first item and remove it from the stack
func (this *Stack) Pop() interface{} {
	this.mu.Lock()
	defer this.mu.Unlock()

	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Add a new item to the stack
func (this *Stack) Push(value interface{}) {
	this.mu.Lock()
	defer this.mu.Unlock()

	n := &node{value, this.top}
	this.top = n
	this.length++
	this.Cond.Signal()
}
