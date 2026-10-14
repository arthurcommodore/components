package main

import (
	"net/http"

	"github.com/arthurcommodore/components/web/example"
	"github.com/arthurcommodore/components/web/middlewares"
)

func main() {
	mux := http.NewServeMux()

	pageHome := middlewares.Chain(middlewares.NoCache)

	// Handler protegido pelo middleware
	mux.Handle("GET /", pageHome(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := example.Page().Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})))

	http.ListenAndServe(":8080", mux)
}
