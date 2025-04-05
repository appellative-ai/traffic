package analytics

import (
	"github.com/behavioral-ai/collective/timeseries"
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

func (l *list) Enqueue(event *timeseries.Event) {
	l.queue.Enque(event)
}

func (l *list) Dequeue() *timeseries.Event {
	item := l.queue.Deque()
	if item == nil {
		return nil
	}
	if event, ok := item.(*timeseries.Event); ok {
		return event
	}
	return nil
}
