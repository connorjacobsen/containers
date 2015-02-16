package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := New()
	size := queue.Size()
	if size != 0 {
		t.Errorf("Expected size to be %s, got %s", 0, size)
	}

	item := queue.Pop()
	if item != nil {
		t.Errorf("Expected value to be %s, got %s", nil, item)
	}

	queue.Push(42)
	size = queue.Size()
	if size != 1 {
		t.Errorf("Expected size to be %s, got %s", 1, size)
	}

	queue.Push(13)
	size = queue.Size()
	if size != 2 {
		t.Errorf("Expected size to be %s, got %s", 2, size)
	}

	queue.Push(7)
	size = queue.Size()
	if size != 3 {
		t.Errorf("Expected size to be %s, got %s", 3, size)
	}

	front := queue.Front()
	if front != 42 {
		t.Errorf("Expected value to be %s, got %s", 42, front)
	}

	back := queue.Back()
	if back != 7 {
		t.Errorf("Expected value to be %s, got %s", 7, back)
	}

	item = queue.Pop()
	if item != 42 {
		t.Errorf("Expected value to be %s, got %s", 42, item)
	}
}
