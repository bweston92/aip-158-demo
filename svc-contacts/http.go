package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math"
	"net"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
)

var (
	//go:embed web
	web embed.FS
)

func startAdminConsole(a *App) error {
	nl, err := net.Listen("tcp", ":3100")
	if err != nil {
		return fmt.Errorf("unable to open port 3100: %w", err)
	}

	webFS, err := fs.Sub(web, "web")
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/api/seed", seedMoreContactsHandler(a))
	mux.Handle("/", http.FileServer(http.FS(webFS)))

	srv := http.Server{
		Handler: mux,
	}

	go srv.Serve(nl)

	return nil
}

func seedMoreContactsHandler(a *App) http.HandlerFunc {
	f := faker.New().Person()

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		w.Header().Add("content-type", "application/json")

		if r.Header.Get("content-type") != "application/json" || r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "invalid request"}`))
			return
		}

		rb, _ := io.ReadAll(r.Body)
		var data struct {
			Count int `json:"count"`
		}
		if err := json.Unmarshal(rb, &data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "json malformed"}`))
			return
		}

		for i := 0; i < int(math.Min(100, float64(data.Count))); i++ {
			c := Contact{
				ID:          uuid.NewString(),
				Forename:    f.FirstName(),
				Surname:     f.LastName(),
				PhoneNumber: f.Contact().Phone,
			}

			if err := a.Add(ctx, &c); err != nil {
				log.Printf("unable to add contact: %s\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "unable to add contact: ` + err.Error() + `"}`))
				return
			}
		}

		w.Write([]byte(`{}`))
	}
}
