package queue

import "container/list"

type Queue struct {
	size int
	list *list.List
}

func New() *Queue { return new(Queue).Init() }

func (q *Queue) Init() *Queue {
	q.list = list.New()
	q.size = 0
	return q
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Front() interface{} {
	return q.list.Front().Value
}

func (q *Queue) Back() interface{} {
	return q.list.Back().Value
}

func (q *Queue) Push(value interface{}) interface{} {
	el := q.list.PushBack(value)
	q.size++
	return el.Value
}

func (q *Queue) Pop() interface{} {
	if q.size == 0 {
		return nil
	}
	el := q.list.Remove(q.list.Front())
	q.size--
	return el
}
