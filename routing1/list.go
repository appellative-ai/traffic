package routing

import (
	"github.com/dustinxie/lockfree"
)

type list struct {
	queue lockfree.Queue
}

func newList() *list {
	b := new(list)
	b.queue = lockfree.NewQueue()
	return b
}

func (l *list) Enqueue(event *event) {
	l.queue.Enque(event)
}

func (l *list) Dequeue() *event {
	item := l.queue.Deque()
	if item == nil {
		return nil
	}
	if v, ok := item.(*event); ok {
		return v
	}
	return nil
}
