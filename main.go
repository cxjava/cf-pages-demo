package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/syumai/workers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
			name := req.URL.Query().Get("name")
			if name == "" {
				name = "Pages Functions"
			}
			fmt.Fprintf(w, "Hello, %s!", name)
		})
		r.Get("/hello2", func(w http.ResponseWriter, req *http.Request) {
			response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
			if err != nil {
				fmt.Print(err.Error())
			}
			responseData, err := io.ReadAll(response.Body)
			if err != nil {
			}
			fmt.Println(string(responseData))

			fmt.Fprintf(w, "Hello, Hello world! plus remote "+string(responseData))
		})
		r.Get("/hello3", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "Hello, Hello, Hello world!")
		})
	})
	workers.Serve(r)
}
