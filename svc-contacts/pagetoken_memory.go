package main

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

type (
	memoryPageToken struct {
		rb  []byte
		li  string
		exp time.Time
	}

	memoryPageTokenProvider struct {
		sync.RWMutex

		t      map[string]*memoryPageToken
		ticker *time.Ticker
	}
)

func newMemoryPageTokenProvider() *memoryPageTokenProvider {
	p := &memoryPageTokenProvider{
		t:      make(map[string]*memoryPageToken),
		ticker: time.NewTicker(time.Minute),
	}
	go p.gc()
	return p
}

// provide a new unique opaque token
func (s *memoryPageTokenProvider) newID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

func (s *memoryPageTokenProvider) gc() {
	for range s.ticker.C {
		s.Lock()
		for k, v := range s.t {
			if time.Now().After(v.exp) {
				delete(s.t, k)
			}
		}
		s.Unlock()
	}
}

func (s *memoryPageTokenProvider) Close() {
	s.Lock()
	defer s.Unlock()

	s.ticker.Stop()
	s.t = nil
}

func (s *memoryPageTokenProvider) Persist(ctx context.Context, filters proto.Message, lastId string) (string, error) {
	s.Lock()
	defer s.Unlock()

	rb, _ := proto.Marshal(filters)

	pt := s.newID()
	s.t[pt] = &memoryPageToken{
		rb: rb,
		li: lastId,
		// depending on the use case, this would usually be around 3 days
		// according to AIP-158.
		exp: time.Now().Add(5 * time.Minute),
	}

	return pt, nil
}

func (s *memoryPageTokenProvider) Find(ctx context.Context, pageToken string, dst proto.Message) (string, error) {
	s.RLock()
	defer s.RUnlock()

	v, ok := s.t[pageToken]
	if !ok {
		return "", errors.New("not found")
	}

	return v.li, proto.Unmarshal(v.rb, dst)
}
