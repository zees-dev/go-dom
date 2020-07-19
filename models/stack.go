package models

import (
	"errors"
	"sync"
)

type Stack struct {
	items   []Item
	rwMutex sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Dump() []Item {
	stack.rwMutex.Lock()
	defer stack.rwMutex.Unlock()

	var items = make([]Item, len(stack.items))
	copy(items, stack.items)
	return items
}

func (stack *Stack) Peek() (Item, error) {
	stack.rwMutex.Lock()
	defer stack.rwMutex.Unlock()

	if len(stack.items) == 0 {
		return Item{}, errors.New("stack is empty...")
	}
	lastItem := stack.items[len(stack.items)-1]
	return lastItem, nil
}

func (stack *Stack) Push(item Item) {
	stack.rwMutex.Lock()
	defer stack.rwMutex.Unlock()

	stack.items = append(stack.items, item)
}

func (stack *Stack) Pop() (Item, error) {
	stack.rwMutex.Lock()
	defer stack.rwMutex.Unlock()

	if len(stack.items) == 0 {
		return Item{}, errors.New("stack is empty...")
	}
	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return lastItem, nil
}

func (stack *Stack) Size() int {
	return len(stack.items)
}

func (stack *Stack) IsEmpty() bool {
	stack.rwMutex.Lock()
	defer stack.rwMutex.Unlock()

	return len(stack.items) == 0
}

func (stack *Stack) Reset() {
	stack.rwMutex.Lock()
	defer stack.rwMutex.Unlock()

	stack.items = []Item{}
}
