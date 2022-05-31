package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
)

var (
	cnfStore = flag.String("contact-store", "inmemory", "Which storage to use for contacts.")
)

type (
	Contact struct {
		ID          string
		Forename    string
		Surname     string
		PhoneNumber string
	}

	Filters struct {
	}

	Store interface {
		List(context.Context, Filters, int32, int32) ([]Contact, error)
		Count(context.Context, Filters) (int32, error)
		Persist(context.Context, *Contact) error
	}
)

func getStoreFromEnvironment() (Store, error) {
	parts := strings.Split(*cnfStore, ":")
	driver := parts[0]

	switch driver {
	case "inmemory":
		return newMemoryStore(), nil
	default:
		return nil, fmt.Errorf("unknown contact store config: %s", driver)
	}
}
