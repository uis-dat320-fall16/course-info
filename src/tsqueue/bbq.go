package tsqueue

import "sync"

// BBQueue is a Bounded Blocking Queue
type BBQueue struct {
	// Synchronization variables
	lock        sync.Mutex
	itemAdded   sync.Cond
	itemRemoved sync.Cond
	// State variables
	items     [max]int
	front     int
	nextEmpty int
}

func (q *BBQueue) insert(item int) {
	q.lock.Lock()
	for q.nextEmpty-q.front == max {
		q.itemRemoved.Wait() // queue is full, wait for remove
	}
	q.items[q.nextEmpty%max] = item
	q.nextEmpty++
	q.itemAdded.Signal()
	q.lock.Unlock()
}

func (q *BBQueue) remove() (item int) {
	q.lock.Lock()
	for q.front == q.nextEmpty {
		q.itemAdded.Wait() // queue is empty, wait for add
	}
	item = q.items[q.front%max]
	q.front++
	q.itemRemoved.Signal()
	q.lock.Unlock()
	return
}
