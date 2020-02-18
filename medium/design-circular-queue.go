package main

import "fmt"

func main() {
	circularQueue := Constructor(3)       // set the size to be 3
	fmt.Println(circularQueue.EnQueue(1)) // return true
	fmt.Println(circularQueue.EnQueue(2)) // return true
	fmt.Println(circularQueue.EnQueue(3)) // return true
	fmt.Println(circularQueue.EnQueue(4)) // return false, the queue is full
	fmt.Println(circularQueue.Rear())     // return 3
	fmt.Println(circularQueue.IsFull())   // return true
	fmt.Println(circularQueue.DeQueue())  // return true
	fmt.Println(circularQueue.EnQueue(4)) // return true
	fmt.Println(circularQueue.Rear())     // return 4
}

type MyCircularQueue struct {
	head  int
	tail  int
	empty bool
	full  bool
	queue []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{0, 0, true, false, make([]int, k)}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.queue[this.tail] = value
	if this.tail < len(this.queue)-1 {
		this.tail++
	} else {
		this.tail = 0
	}
	if this.tail == this.head {
		this.full = true
	}
	if this.empty {
		this.empty = false
	}
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	if this.head < len(this.queue)-1 {
		this.head++
	} else {
		this.head = 0
	}
	if this.tail == this.head {
		this.empty = true
	}
	if this.full {
		this.full = false
	}
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.empty {
		return -1
	}
	return this.queue[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.empty {
		return -1
	}
	if this.tail == 0 {
		return this.queue[len(this.queue)-1]
	}
	return this.queue[this.tail-1]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.empty
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.full
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
