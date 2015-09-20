package tsqueue

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func ExampleChQueue() {
	q := NewChQueue()
	// q.insert(10)
	q.insert(10)
	q.insert(20)
	q.insert(30)
	q.insert(40)
	for i := 0; i < 4; i++ {
		e := q.remove()
		fmt.Print(e, " ")
	}
	//Output: 10 20 30 40
}

func TestChQueue(t *testing.T) {
	var q [threads]*ChQueue
	for i := 0; i < threads; i++ {
		q[i] = NewChQueue()
		wg.Add(1)
		go insertSome(q[i], i)
	}
	time.Sleep(10 * time.Millisecond)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go removeSome(q[i], i)
	}
	wg.Wait()
}

func insertSome(q *ChQueue, id int) {
	for i := 0; i < 20; i++ {
		q.insert(i)
	}
	wg.Done()
}

func removeSome(q *ChQueue, id int) {
	for i := 0; i < 20; i++ {
		item := q.remove()
		fmt.Printf("Removed item %d (thread %d)\n", item, id)
		runtime.Gosched()
	}
	wg.Done()
}
