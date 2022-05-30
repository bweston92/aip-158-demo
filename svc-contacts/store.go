package main

import "context"

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
