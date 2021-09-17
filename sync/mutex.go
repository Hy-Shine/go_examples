package main

import (
	"fmt"
	"sync"
	"time"
)

type mutexT struct {
	safe sync.Mutex
	v    map[string]int
}

func (m *mutexT) intCount(key string) {
	m.safe.Lock()
	m.v[key]++
	m.safe.Unlock()
}

func (m *mutexT) getKey(key string) int {
	m.safe.Lock()
	defer m.safe.Unlock()
	return m.v[key]
}

func mutexHandle() {
	m := &mutexT{
		v: make(map[string]int),
	}
	for i := 0; i < 500; i++ {
		go m.intCount("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(m.getKey("somekey"))
}
