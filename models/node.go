package models

import (
	"fmt"
	"sync"
)

type Node struct {
	tag      string
	id       string
	src      string
	alt      string
	text     string
	class    string
	children []*Node
}

// Equals comparison operator for Nodes
func (node *Node) Equals(other *Node) bool {
	primitivesMatch := node.tag == other.tag &&
		node.id == other.id &&
		node.src == other.src &&
		node.alt == other.alt &&
		node.text == other.text &&
		node.class == other.class &&
		len(node.children) == len(other.children)

	if primitivesMatch == false {
		return false
	}

	if len(node.children) == len(other.children) {
		for i := range node.children {
			if !node.children[i].Equals(other.children[i]) {
				return false
			}
		}
	}

	return true
}

func (node *Node) String() string {
	return fmt.Sprintf("{ id: %s, tag: %s, src: %s, alt: %s, text: %s, class: %s, children: %q }",
		node.id,
		node.tag,
		node.src,
		node.alt,
		node.text,
		node.class,
		node.children,
	)
}

// Item - Go doesn't yet support generics hence we pre-declare stack/queue type here
type Item = Node

// GetElementByID is optimised node traversal to find element by specified ID using of go-routines and context
func (node *Node) GetElementByID(id string) (*Node, error) {
	var wg sync.WaitGroup
	wg.Add(1)
	nodeCh := make(chan *Node)

	var getNodeByID func(*Node, *sync.WaitGroup)
	getNodeByID = func(n *Node, wg *sync.WaitGroup) {
		defer wg.Done()

		if n.id == id {
			nodeCh <- n
		}
		for _, v := range n.children {
			wg.Add(1)
			go getNodeByID(v, wg)
		}
	}
	getNodeByID(node, &wg)

	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(nodeCh)
	}(&wg)

	nodeWithID, found := <-nodeCh
	if !found {
		return nil, fmt.Errorf("element with id %s not found", id)
	}

	// Use context to cancel lookup once a single goroutine has found node by ID
	// For for - select with context
	return nodeWithID, nil
}

// GetElementByIDBFS is Breadth first search implementation of node traversal to retrieve node by ID
func (node *Node) GetElementByIDBFS(id string) (*Node, error) {
	queue := NewQueue()
	queue.EnQueue(*node)
	for !queue.IsEmpty() {
		currentNode, err := queue.DeQueue()
		if err != nil {
			break
		}
		if currentNode.id == id {
			return &currentNode, nil
		}
		for _, v := range currentNode.children {
			queue.EnQueue(*v)
		}
	}
	return nil, fmt.Errorf("element with id %s not found", id)
}

// GetElementByIDDFS is Depth first search implementation of node traversal to retrieve node by ID
func (node *Node) GetElementByIDDFS(id string) (*Node, error) {
	stack := NewStack()
	stack.Push(*node)
	for !stack.IsEmpty() {
		currentNode, err := stack.Pop()
		if err != nil {
			break
		}
		if currentNode.id == id {
			return &currentNode, nil
		}
		for _, v := range currentNode.children {
			stack.Push(*v)
		}
	}
	return nil, fmt.Errorf("element with id %s not found", id)
}

// GetExampleDom returns a mock/example DOM
func GetExampleDom() *Node {
	image := Node{
		tag: "img",
		src: "http://example.com/logo.svg",
		alt: "Example's Logo",
	}

	p := Node{
		tag:      "p",
		text:     "And this is some text in a paragraph. And next to it there's an image.",
		children: []*Node{&image},
	}

	span := Node{
		tag:  "span",
		id:   "copyright",
		text: "2019 ; Zees-Dev",
	}

	div := Node{
		tag:      "div",
		class:    "footer",
		text:     "This is the footer of the page.",
		children: []*Node{&span},
	}

	h1 := Node{
		tag:  "h1",
		text: "This is a H1",
	}

	body := Node{
		tag:      "body",
		children: []*Node{&h1, &p, &div},
	}

	html := Node{
		tag:      "html",
		children: []*Node{&body},
	}

	return &html
}
