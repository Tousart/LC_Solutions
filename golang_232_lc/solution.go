package main

import "fmt"

type MyQueue struct {
	put  []int
	take []int
}

func Constructor() MyQueue {
	return MyQueue{
		put:  []int{},
		take: []int{},
	}
}

func (mq *MyQueue) Push(x int) {
	mq.put = append(mq.put, x)
}

func (mq *MyQueue) Pop() int {
	var x int

	if len(mq.take) > 0 {
		x = mq.take[len(mq.take)-1]
	} else {
		x = distillation(mq)
	}
	mq.take = mq.take[:len(mq.take)-1]

	return x
}

func (mq *MyQueue) Peek() int {
	var x int

	if len(mq.take) > 0 {
		x = mq.take[len(mq.take)-1]
	} else {
		x = distillation(mq)
	}

	return x
}

func (mq *MyQueue) Empty() bool {
	return len(mq.put) == 0 && len(mq.take) == 0
}

func distillation(queue *MyQueue) int {
	var x int

	lenPut := len(queue.put)
	for lenPut > 0 {
		lenPut--
		x = queue.put[lenPut]
		queue.take = append(queue.take, x)
		queue.put = queue.put[:lenPut]
	}

	return x
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Empty())
}
