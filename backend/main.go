package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/api/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("One of the things Ford Prefect had always found hardest to understand about humans was their habit of continually stating and repeating the very very obvious."))
	})

	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
