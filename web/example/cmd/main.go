package main

import (
	"net/http"
	"strings"

	"github.com/arthurcommodore/components/language"
	"github.com/arthurcommodore/components/web/example"
	"github.com/arthurcommodore/components/web/middlewares"
)

func main() {
	mux := http.NewServeMux()

	pageHome := middlewares.Chain(middlewares.NoCache)

	mux.Handle("GET /", pageHome(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := language.NewLanguage("lab/languages", "pt")

		currentLanguage := strings.TrimPrefix(r.URL.Path, "/")

		if currentLanguage == "" {
			currentLanguage = lang.DefaultLanguage
		}

		lang.DefaultLanguage = currentLanguage
		err := example.Page(lang).Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})))

	http.ListenAndServe(":8080", mux)
}
