package main

import (
	"context"
	"errors"
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

func (s *memoryStore) Persist(_ context.Context, c *Contact) error {
	s.m.Lock()
	defer s.m.Unlock()

	if c.ID == "" {
		return errors.New("invalid identifier")
	}

	if _, ok := s.id[c.ID]; ok {
		return errors.New("conflict")
	}

	s.id[c.ID] = len(s.entries)
	s.entries = append(s.entries, *c)

	return nil
}
