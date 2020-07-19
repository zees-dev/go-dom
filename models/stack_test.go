package models

import "testing"

func TestStack(t *testing.T) {

	t.Run("peek stack", func(t *testing.T) {
		stack := NewStack()
		stack.Push(Node{id: "1"})
		stack.Push(Node{id: "2"})
		got, err := stack.Peek()
		if err != nil {
			t.Errorf("error %v", err)
		}
		want := Node{id: "2"}

		if !got.Equals(&want) {
			t.Errorf("got %v want %v", got, want)
		}

		gotSize := stack.Size()
		wantSize := 2
		if gotSize != wantSize {
			t.Errorf("gotSize %v wantSize %v", got, want)
		}
	})

	t.Run("pop stack", func(t *testing.T) {
		stack := NewStack()
		stack.Push(Node{id: "1"})
		stack.Push(Node{id: "2"})
		got, err := stack.Pop()
		if err != nil {
			t.Errorf("error %v", err)
		}
		want := Node{id: "2"}

		if !got.Equals(&want) {
			t.Errorf("got %v want %v", got, want)
		}

		gotSize := stack.Size()
		wantSize := 1
		if gotSize != wantSize {
			t.Errorf("gotSize %v wantSize %v", got, want)
		}
	})

	t.Run("empty stack", func(t *testing.T) {
		stack := NewStack()
		stack.Push(Node{id: "1"})
		stack.Push(Node{id: "2"})
		stack.Reset()

		got := stack.IsEmpty()
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
