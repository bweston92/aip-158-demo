package main

import (
	"context"
	"testing"

	"github.com/bweston92/aip-158-demo/svc-contacts/api"
	"google.golang.org/protobuf/proto"
)

func TestMemoryPageTokenProvider(t *testing.T) {
	ctx := context.Background()
	sut := newMemoryPageTokenProvider()

	expected := &api.ListContactsRequest_Filters{
		Forename:    "Tester",
		PhoneNumber: "011",
	}

	pt, err := sut.Persist(ctx, expected, "abcd1234")
	if err != nil {
		t.Fatal("unable to get a page token")
	}

	actual := &api.ListContactsRequest_Filters{}
	lastID, err := sut.Find(ctx, pt, actual)
	if err != nil {
		t.Fatal("page token was not resolved", err)
	}

	if lastID != "abcd1234" {
		t.Errorf("last identifier was incorrect, expected abcd1234 got %s", lastID)
	}

	if !proto.Equal(actual, expected) {
		t.Errorf("filters mismatch:\n\texpected: %+v\n\tactual: %+v", expected, actual)
	}
}
