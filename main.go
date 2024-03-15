package main

import (
	"fmt"
)

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}

}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(s string) {
	node := &Node{}

	if val, ok := c.Hash[s]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: s}
	}
	c.Add(node)
	c.Hash[s] = node
}

func (c *Cache) Remove(node *Node) *Node {
	fmt.Printf("Remove node %s", node.Val)
	fmt.Println()
	left := node.Left
	right := node.Right

	left.Right = right
	right.Left = left

	c.Queue.Length -= 1
	delete(c.Hash, node.Val)
	return node
}

func (c *Cache) Add(node *Node) {
	fmt.Printf("Add node %s", node.Val)
	fmt.Println()
	temp := c.Queue.Head.Right

	c.Queue.Head.Right = node
	node.Right = temp
	node.Left = c.Queue.Head
	temp.Left = node

	c.Queue.Length += 1

	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}

}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {

	node := q.Head.Right

	fmt.Printf(" %d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Printf("]")
}

type Hash map[string]*Node

func main() {
	fmt.Println("Welcome to LRU Cache")
	cache := NewCache()
	for _, word := range []string{"hello", "my", "name", "is", "pritish", "is", "tree"} {
		cache.Check(word)
		cache.Display()
		fmt.Println()
	}
}
