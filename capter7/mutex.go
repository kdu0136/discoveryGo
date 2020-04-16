package capter7

import "sync"

func MutexMain() {
}

type Accessor struct {
	R *Response
	L *sync.Mutex
}

func (acc *Accessor) Use() {
	// do something
	acc.L.Lock()
	// use acc.R
	acc.L.Unlock()
	// do something else
}

type ConcurrentMap struct {
	M map[string]string
	L *sync.RWMutex
}

func (m ConcurrentMap) Get(key string) string {
	m.L.RLock()
	defer m.L.RUnlock()
	return m.M[key]
}

func (m ConcurrentMap) Set(key, value string) {
	m.L.Lock()
	defer m.L.Unlock()
	m.M[key] = value
}
