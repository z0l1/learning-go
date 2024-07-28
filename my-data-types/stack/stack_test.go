package stack

import "testing"

func TestNew(t *testing.T) {
	stack := New[string]()
	if stack.data == nil {
		t.Fatal("stack data is nil")
	}
}

func TestGood(t *testing.T) {
}
