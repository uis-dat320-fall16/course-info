package tsqueue

import (
	"sync"
	"time"
)

type TSQueue struct {
	// Synchronization variables
	lock *sync.Mutex
	// State variables
	items     [max]int
	front     int
	nextEmpty int
}

func NewTSQueue() *TSQueue {
	return &TSQueue{lock: &sync.Mutex{}}
}

func (q *TSQueue) insertX(item int) {
	for !q.tryInsert(item) {
		time.Sleep(1 * time.Millisecond)
	}
}

func (q *TSQueue) tryInsert(item int) (success bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.nextEmpty-q.front < max {
		q.items[q.nextEmpty%max] = item
		q.nextEmpty++
		success = true
	}
	return
}

func (q *TSQueue) tryRemove(item *int) (success bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.front < q.nextEmpty {
		*item = q.items[q.front%max]
		q.front++
		success = true
	}
	return
}
