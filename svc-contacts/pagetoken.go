package main

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type (
	PageTokenProvider interface {
		Persist(ctx context.Context, filters proto.Message, lastId string) (string, error)

		// Find a page token, populate the `dst` with the filters and return a string
		// that identifies the last result in the previous result set.
		Find(ctx context.Context, pageToken string, dst proto.Message) (string, error)
	}
)
