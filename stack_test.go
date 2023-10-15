package stack

import (
	"testing"
)

func Test(t *testing.T) {

	s := New()

	if s.Len() != 0 {
		t.Errorf("Len empty stack should be 0")
	}

	s.Push(1)

	if s.Len() != 1 {
		t.Errorf("Top item of stack of 1 should be 1")
	}

	if s.Peek().(int) != 1 {
		t.Errorf("Top item should be 1 when peeked")
	}

	if s.Pop().(int) != 1 {
		t.Errorf("Top item should be 1 when popped")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty after a pop")
	}

	s.Push(1)
	s.Push(2)

	if s.Len() != 2 {
		t.Errorf("Len should be 2")
	}

	if s.Peek().(int) != 2 {
		t.Errorf("Top of the stack should be 2")
	}

}
