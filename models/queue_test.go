package models

import "testing"

func TestQueue(t *testing.T) {

	t.Run("peek queue", func(t *testing.T) {
		queue := NewQueue()
		queue.EnQueue(Node{id: "1"})
		queue.EnQueue(Node{id: "2"})
		got, err := queue.Peek()
		if err != nil {
			t.Errorf("error %v", err)
		}
		want := Node{id: "1"}

		if !got.Equals(&want) {
			t.Errorf("got %v want %v", got, want)
		}

		gotSize := queue.Size()
		wantSize := 2
		if gotSize != wantSize {
			t.Errorf("gotSize %v wantSize %v", got, want)
		}
	})

	t.Run("dequeue queue", func(t *testing.T) {
		queue := NewQueue()
		queue.EnQueue(Node{id: "1"})
		queue.EnQueue(Node{id: "2"})
		got, err := queue.DeQueue()
		if err != nil {
			t.Errorf("error %v", err)
		}
		want := Node{id: "1"}

		if !got.Equals(&want) {
			t.Errorf("got %v want %v", got, want)
		}

		gotSize := queue.Size()
		wantSize := 1
		if gotSize != wantSize {
			t.Errorf("gotSize %v wantSize %v", got, want)
		}
	})

	t.Run("empty queue", func(t *testing.T) {
		queue := NewQueue()
		queue.EnQueue(Node{id: "1"})
		queue.EnQueue(Node{id: "1"})
		queue.Reset()

		got := queue.IsEmpty()
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
