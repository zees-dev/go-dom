package models

import (
	"errors"
	"sync"
)

// Go doesn't yet support generics hence we pre-declare stack type

type Queue struct {
	items   []Item
	rwMutex sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{}
}

func (queue *Queue) Dump() []Item {
	queue.rwMutex.Lock()
	defer queue.rwMutex.Unlock()

	var items = make([]Item, len(queue.items))
	copy(items, queue.items)
	return items
}

func (queue *Queue) Peek() (Item, error) {
	queue.rwMutex.Lock()
	defer queue.rwMutex.Unlock()

	if len(queue.items) == 0 {
		return Item{}, errors.New("stack is empty...")
	}
	firstItem := queue.items[0]
	return firstItem, nil
}

func (queue *Queue) EnQueue(item Item) {
	queue.rwMutex.Lock()
	defer queue.rwMutex.Unlock()

	queue.items = append(queue.items, item)
}

func (queue *Queue) DeQueue() (Item, error) {
	queue.rwMutex.Lock()
	defer queue.rwMutex.Unlock()

	if len(queue.items) == 0 {
		return Item{}, errors.New("queue is empty...")
	}
	firstItem := queue.items[0]
	queue.items = queue.items[1:len(queue.items)]
	return firstItem, nil
}

func (queue *Queue) Size() int {
	return len(queue.items)
}

func (queue *Queue) IsEmpty() bool {
	queue.rwMutex.Lock()
	defer queue.rwMutex.Unlock()

	return len(queue.items) == 0
}

func (queue *Queue) Reset() {
	queue.rwMutex.Lock()
	defer queue.rwMutex.Unlock()

	queue.items = []Item{}
}
