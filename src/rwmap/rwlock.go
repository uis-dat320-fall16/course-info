package main

import "sync"

//RWLock some more
type RWLock struct {
	// synchronziation variables
	mu      *sync.Mutex
	readGo  *sync.Cond
	writeGo *sync.Cond

	// state variables
	activeReaders, waitingReaders int
	activeWriters, waitingWriters int
}

// NewRWLock returns an RWLock
func NewRWLock() *RWLock {
	rwl := &RWLock{}
	rwl.mu = &sync.Mutex{}
	rwl.readGo = sync.NewCond(rwl.mu)
	rwl.writeGo = sync.NewCond(rwl.mu)
	return rwl
}

func (rw *RWLock) startRead() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.waitingReaders++
	for rw.readShouldWait() {
		rw.readGo.Wait()
	}
	rw.waitingReaders--
	rw.activeReaders++
}

func (rw *RWLock) doneRead() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.activeReaders--
	if rw.activeReaders == 0 && rw.waitingWriters > 0 {
		rw.writeGo.Signal()
	}
}

func (rw *RWLock) startWrite() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.waitingWriters++
	for rw.writeShouldWait() {
		rw.writeGo.Wait()
	}
	rw.waitingWriters--
	rw.activeWriters++
}

func (rw *RWLock) doneWrite() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.activeWriters--
	// assert activeWriters == 0
	if rw.waitingWriters > 0 {
		rw.writeGo.Signal()
	} else {
		// here there may be waiting readers
		rw.readGo.Broadcast()
	}
}

func (rw *RWLock) readShouldWait() bool {
	return rw.activeWriters > 0 || rw.waitingWriters > 0
}

func (rw *RWLock) writeShouldWait() bool {
	return rw.activeWriters > 0 || rw.activeReaders > 0
}
