package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/syumai/workers"
)

func main() {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
			name := req.URL.Query().Get("name")
			if name == "" {
				name = "Pages Functions"
			}
			fmt.Fprintf(w, "Hello, %s!", name)
		})
		r.Get("/hello2", func(w http.ResponseWriter, req *http.Request) {
			resp, err := http.DefaultClient.Get(`https://cf-pages-demo-cxjava.pages.dev/api/hello`)
			if err != nil {
				fmt.Fprintf(w, "Hello, Hello world! plus "+err.Error())
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(req.Body)
			fmt.Fprintf(w, "Hello, Hello world! plus remote "+string(body))
		})
		r.Get("/hello3", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "Hello, Hello, Hello world!")
		})
	})
	workers.Serve(r)
}
