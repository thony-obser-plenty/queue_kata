package main

import "fmt"

func main() {
	queue := NewQueue()

	queue.Enqueue(5)
	queue.Enqueue(10)
	queue.Enqueue(15)

	val := queue.Peek()
	fmt.Println(val) // Output: 5

	val = queue.Dequeue()
	fmt.Println(val) // Output: 5

	isEmpty := queue.IsEmpty()
	fmt.Println(isEmpty) // Output: false
}

type Queue struct {
	data []any
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(element any) {
	q.data = append(q.data, element)
}

func (q *Queue) Dequeue() any {
	if q.IsEmpty() {
		return false
	}

	entry := q.data[0]
	q.data = q.data[1:]

	return entry
}

func (q *Queue) Peek() any {
	if q.IsEmpty() {
		return false
	}

	entry := q.data[0]

	return entry
}

func (q *Queue) IsEmpty() bool {
	isEmpty := len(q.data) == 0
	return isEmpty
}
