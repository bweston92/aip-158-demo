package main

import "context"

type App struct {
	contacts Store
}

func newApp(s Store) *App {
	return &App{
		contacts: s,
	}
}

func (a *App) Add(ctx context.Context, v *Contact) error {
	return a.contacts.Persist(ctx, v)
}
