package bst

import "testing"

func TestBST(t *testing.T) {
	bst := New()
	size := bst.Size()
	if size != 0 {
		t.Errorf("Expected size to be %s, got %s", 0, size)
	}

	root := bst.Root()
	if root != nil {
		t.Errorf("Expected root to be nil, got %s", root)
	}

	item := bst.Insert(27, "foo")
	if item != "foo" {
		t.Errorf("Expected return value to be %s, got %s", "foo", item)
	}
	size = bst.Size()
	if size != 1 {
		t.Errorf("Expected size to be %s, got %s", 1, size)
	}
	node := bst.Root()
	if node.Key != 27 {
		t.Errorf("Expected value to be %s, got %s", 27, node.Key)
	}
	if node.Value != "foo" {
		t.Errorf("Expected return value to be %s, got %s", "foo", item)
	}
}
