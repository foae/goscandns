package main

import "sync"

type registry struct {
	l *sync.RWMutex
	m map[string]int
}

func newRegistry() *registry {
	return &registry{
		l: new(sync.RWMutex),
		m: make(map[string]int),
	}
}

func (r *registry) add(s string) {
	r.l.Lock()
	r.m[s]++
	r.l.Unlock()
}

func (r *registry) retrieve() map[string]int {
	r.l.RLock()
	defer r.l.RUnlock()
	return r.m
}

func (r *registry) len() int {
	r.l.RLock()
	defer r.l.RUnlock()
	return len(r.m)
}

func (r *registry) String() string {
	return string(r.len())
}
