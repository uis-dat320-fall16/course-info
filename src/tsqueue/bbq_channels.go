package tsqueue

// ChQueue is a Bounded Blocking Queue
type ChQueue struct {
	// State variables
	items chan int
}

func NewChQueue() *ChQueue {
	return &ChQueue{items: make(chan int, max)}
}

func (q *ChQueue) insert(item int) {
	q.items <- item
}

func (q *ChQueue) remove() (item int) {
	return <-q.items
}
