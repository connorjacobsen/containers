package stack

import "testing"

func TestStack(t *testing.T) {
	stack := New()
	if stack.Peek() != nil {
		t.Errorf("Stack initialization error.")
	}

	if stack.Size() != 0 {
		t.Errorf("Expected stack size to be %s, was %s", 0, stack.Size())
	}

	item := stack.Push(42)
	if item != 42 {
		t.Errorf("Expected return value to be %s, got %s", 42, item)
	}
	if stack.Size() != 1 {
		t.Errorf("Expected stack Size to be %s, got %s", 1, stack.Size())
	}

	item = stack.Push(13)
	if item != 13 {
		t.Errorf("Expected return value to be %s, got %s", 13, item)
	}
	if stack.Size() != 2 {
		t.Errorf("Expected stack Size to be %s, got %s", 2, stack.Size())
	}

	item = stack.Peek()
	if item != 13 {
		t.Errorf("Expected return value to be %s, got %s", 13, item)
	}
	if stack.Size() != 2 {
		t.Errorf("Expected stack Size to be %s, got %s", 2, stack.Size())
	}

	item = stack.Pop()
	if item != 13 {
		t.Errorf("Expected return value to be %s, got %s", 13, item)
	}
	if stack.Size() != 1 {
		t.Errorf("Expected stack Size to be %s, got %s", 1, stack.Size())
	}
}
