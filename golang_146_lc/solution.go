package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	val  int
	key  int
}

type LRUCache struct {
	valNode  map[int]*Node
	root     *Node
	count    int
	capacity int
}

func Constructor(capacity int) LRUCache {
	valNode := make(map[int]*Node)
	root := Node{val: -1, key: -1}
	root.next = &root
	root.prev = &root
	count := 0
	return LRUCache{
		valNode:  valNode,
		root:     &root,
		count:    count,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node := this.valNode[key]
	if node != nil {
		del(node)
		add(this.root, node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node := this.valNode[key]

	if node != nil {
		node.val = value
		del(node)
		add(this.root, node)
	} else {
		node = &Node{val: value, key: key}
		if this.count < this.capacity {
			add(this.root, node)
			this.count++
		} else {
			oldNode := this.root.next
			del(oldNode)
			delete(this.valNode, oldNode.key)
			add(this.root, node)
		}
		this.valNode[key] = node
	}
}

func add(root, node *Node) {
	prevv := root.prev
	root.prev = node
	prevv.next = node
	node.prev = prevv
	node.next = root
}

func del(node *Node) {
	prevv := node.prev
	nextt := node.next
	prevv.next = nextt
	nextt.prev = prevv
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
}
