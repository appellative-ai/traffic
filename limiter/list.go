package limiter

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

func (l *list) enqueue(event *event) {
	l.queue.Enque(event)
}

func (l *list) dequeue() *event {
	item := l.queue.Deque()
	if item == nil {
		return nil
	}
	if v, ok := item.(*event); ok {
		return v
	}
	return nil
}

func (l *list) empty() {
	for item := l.queue.Deque(); item != nil; {
	}
}
