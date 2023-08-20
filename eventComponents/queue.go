package eventComponents

import (
	"sync"
)

type Queue struct {
	items []Event
	lock  sync.Mutex
}

func (q *Queue) Enqueue(item Event) {
	q.lock.Lock()         // Lock the queue so that no other threads can enqueue events at the same time
	defer q.lock.Unlock() // Unlock the queue when the function ends
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() Event {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.items) == 0 {
		return Event{}
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
