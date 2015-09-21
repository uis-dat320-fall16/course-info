package main

import (
	"fmt"
	"time"
)

type rwlockmap struct {
	ma map[string]string
	rw *RWLock
}

func main() {
	smap := &rwlockmap{make(map[string]string), NewRWLock()}

	for i := 0; i < 1000; i++ {
		go func() {
			for {
				x := smap.lookup("hein")
				fmt.Println(x)
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func(j int) {
			val := fmt.Sprintf("meling %d", j)
			for {
				smap.insert("hein", val)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	ch := make(chan bool)
	<-ch
}

func (m *rwlockmap) lookup(key string) string {
	m.rw.startRead()
	defer m.rw.doneRead()
	return m.ma[key]
}

func (m *rwlockmap) insert(key, value string) {
	m.rw.startWrite()
	defer m.rw.doneWrite()
	m.ma[key] = value
}
