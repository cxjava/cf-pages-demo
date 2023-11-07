package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jpillora/installer/handler"
	"github.com/syumai/workers"
)

func main() {
	r := chi.NewRouter()
	c := handler.DefaultConfig
	h := &handler.Handler{Config: c}
	r.Route("/api", func(r chi.Router) {
		// https://github.com/jpillora/installer
		r.Get("/install", h.ServeHTTP)
		
		r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
			name := req.URL.Query().Get("name")
			if name == "" {
				name = "Pages Functions"
			}
			fmt.Fprintf(w, "Hello, %s!", name)
		})
		r.Get("/hello2", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "Hello, Hello world!")
		})
		r.Get("/hello3", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "Hello, Hello, Hello world!")
		})
	})
	workers.Serve(r)
}
