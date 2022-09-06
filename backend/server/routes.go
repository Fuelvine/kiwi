package server

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
