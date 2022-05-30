package main

import (
	"context"
	"sync"
)

var (
	_ Store = (*memoryStore)(nil)
)

type memoryStore struct {
	m       *sync.RWMutex
	id      map[string]int
	entries []Contact
}

func newMemoryStore() *memoryStore {
	return &memoryStore{
		m:       &sync.RWMutex{},
		id:      map[string]int{},
		entries: make([]Contact, 0, 100),
	}
}

func (s *memoryStore) List(context.Context, Filters, int32, int32) ([]Contact, error) {
	return nil, nil
}

func (s *memoryStore) Count(context.Context, Filters) (int32, error) {
	return 0, nil
}

func (s *memoryStore) Persist(context.Context, *Contact) error {
	return nil
}
