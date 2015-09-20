package tsqueue

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestTSQueue(t *testing.T) {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(1)
	var q [threads]*TSQueue
	for i := 0; i < threads; i++ {
		q[i] = NewTSQueue()
		wg.Add(1)
		go putSome(q[i], i)
	}
	time.Sleep(10 * time.Millisecond)
	// wg.Wait()

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go testRemoval(q[i], i)
	}
	wg.Wait()
}

func putSome(q *TSQueue, id int) {
	for i := 0; i < 20; i++ {
		q.insertX(i)
		// fmt.Printf("inserted %d (thread %d)\n", i, id)
		// if !q.tryInsert(i) {
		// 	fmt.Printf("Insert item %d failed (thread %d)\n", i, id)
		// }
	}
	fmt.Printf("done %d\n", id)
	wg.Done()
}

func testRemoval(q *TSQueue, id int) {
	var item int
	for i := 0; i < 20; i++ {
		if q.tryRemove(&item) {
			// 	fmt.Printf("Removed item %d (thread %d)\n", item, id)
			// } else {
			fmt.Printf("Nothing there to remove (thread %d)\n", id)
			runtime.Gosched()
		}
	}
	wg.Done()
}
