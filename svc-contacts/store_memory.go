package main

import (
	"context"
	"errors"
	"strings"
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
		m: &sync.RWMutex{},
		// store the index relative from the end of the slice
		id: map[string]int{},
		// entries stored in newest first order
		entries: make([]Contact, 0, 100),
	}
}

func (s *memoryStore) List(_ context.Context, f Filters, offset int32, limit int32) ([]Contact, error) {
	doneAfterID := f.AfterID == ""
	out := make([]Contact, 0, limit)

	for i := len(s.entries) - 1; i >= 0; i-- {
		c := s.entries[i]
		if !doneAfterID {
			if c.ID == f.AfterID {
				doneAfterID = true
			}
			continue
		}

		if offset > 0 {
			offset--
			continue
		}

		if s.match(&c, &f) {
			out = append(out, c)
		}

		if len(out) == int(limit) {
			break
		}
	}

	return out, nil
}

func (s *memoryStore) match(c *Contact, f *Filters) bool {
	if f.Forename != "" {
		v := strings.ToLower(c.Forename)
		m := strings.ToLower(f.Forename)
		if !strings.Contains(v, m) {
			return false
		}
	}

	if f.PhoneNumber != "" {
		v := strings.ReplaceAll(c.PhoneNumber, "-", "")
		v = strings.ReplaceAll(v, " ", "")

		m := strings.ReplaceAll(f.PhoneNumber, "-", "")
		m = strings.ReplaceAll(v, " ", "")

		if !strings.Contains(v, m) {
			return false
		}
	}

	return true
}

func (s *memoryStore) Count(_ context.Context, f Filters) (int32, error) {
	var t int32 = 0
	for _, v := range s.entries {
		if s.match(&v, &f) {
			t++
		}
	}
	return t, nil
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
	s.entries = append([]Contact{*c}, s.entries...)

	return nil
}
