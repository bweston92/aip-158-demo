package main

import (
	"context"
	"testing"
)

func TestMemoryStore(t *testing.T) {
	ctx := context.Background()
	s := newMemoryStore()

	cs := []Contact{
		{
			ID:          "96794026-5e44-443b-88c7-1237e6a75b93",
			Forename:    "AAA",
			Surname:     "BBB",
			PhoneNumber: "0116 2222 222",
		},
		{
			ID:          "88d8c1f2-219c-4300-8b15-fe460fe34e74",
			Forename:    "YC",
			Surname:     "C",
			PhoneNumber: "0112 3333 333",
		},
		{
			ID:          "f5f2c9ee-0e01-47fe-b385-1e79a2bae4eb",
			Forename:    "YE",
			Surname:     "E",
			PhoneNumber: "0116 1111 111",
		},
		{
			ID:          "6ed72d67-b1b3-47a1-99ed-576b6c82b32a",
			Forename:    "YG",
			Surname:     "G",
			PhoneNumber: "0116 3333 333",
		},
		{
			ID:          "3b1d2d40-96cc-4039-8045-175b03e673db",
			Forename:    "YI",
			Surname:     "I",
			PhoneNumber: "0115 2222 222",
		},
		{
			ID:          "cbfe0cdc-cdbd-4a91-bea0-89ac8eb510f6",
			Forename:    "YJ",
			Surname:     "Y",
			PhoneNumber: "0113 2222 222",
		},
	}

	for _, c := range cs {
		if err := s.Persist(ctx, &c); err != nil {
			t.Fatal("unable to initialise in memory store", err)
		}
	}

	{ // test total
		total, _ := s.Count(ctx, Filters{})
		if v := int32(len(cs)); v != total {
			t.Errorf("expected total to be %d got %d", v, total)
		}
	}

	{ // test total with a known result
		total, _ := s.Count(ctx, Filters{
			Forename: "a",
		})
		if 1 != total {
			t.Errorf("expected total to be 1 got %d", total)
		}
	}

	{ // get a list of results
		results, _ := s.List(ctx, Filters{
			Forename: "y",
		}, 0, 2)

		if len(results) != 2 {
			t.Error("expected 2 results")
		} else {
			if results[0].ID != "88d8c1f2-219c-4300-8b15-fe460fe34e74" {
				t.Error("invalid identifier on result offset 0")
			}

			if results[1].ID != "f5f2c9ee-0e01-47fe-b385-1e79a2bae4eb" {
				t.Error("invalid identifier on result offset 1")
			}
		}
	}

	{ // get a list of results after "f5f2c9ee-0e01-47fe-b385-1e79a2bae4eb"
		results, _ := s.List(ctx, Filters{
			AfterID:  "f5f2c9ee-0e01-47fe-b385-1e79a2bae4eb",
			Forename: "y",
		}, 0, 2)

		if len(results) != 2 {
			t.Error("expected 2 results")
		} else {
			if results[0].ID != "6ed72d67-b1b3-47a1-99ed-576b6c82b32a" {
				t.Error("invalid identifier on result offset 0")
			}

			if results[1].ID != "3b1d2d40-96cc-4039-8045-175b03e673db" {
				t.Error("invalid identifier on result offset 1")
			}
		}
	}

	{ // get a list of results after "f5f2c9ee-0e01-47fe-b385-1e79a2bae4eb" with offset
		results, _ := s.List(ctx, Filters{
			AfterID:  "f5f2c9ee-0e01-47fe-b385-1e79a2bae4eb",
			Forename: "y",
		}, 2, 1)

		if len(results) != 1 {
			t.Error("expected 1 result")
		} else {
			if results[0].ID != "cbfe0cdc-cdbd-4a91-bea0-89ac8eb510f6" {
				t.Error("invalid identifier on result offset 0")
			}
		}
	}

	{ // no more results should be provided
		results, _ := s.List(ctx, Filters{
			AfterID:  "f5f2c9ee-0e01-47fe-b385-1e79a2bae4eb",
			Forename: "y",
		}, 6, 1)

		if len(results) != 0 {
			t.Error("expected 0 result")
		}
	}

	{ // nothing should match, 0 results
		results, _ := s.List(ctx, Filters{
			Forename: "l",
		}, 2, 1)

		if len(results) != 0 {
			t.Error("expected 0 results")
		}
	}
}
