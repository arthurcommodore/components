package main

import (
	"net/http"

	"github.com/arthurcommodore/components/language"
	"github.com/arthurcommodore/components/web/example"
	"github.com/arthurcommodore/components/web/middlewares"
)

func main() {
	mux := http.NewServeMux()

	pageHome := middlewares.Chain(middlewares.NoCache)

	mux.Handle("GET /{lang}", pageHome(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := language.NewLanguage("lab/languages", "pt")

		currentLanguage := r.PathValue("lang")

		if currentLanguage == "" {
			currentLanguage = lang.DefaultLanguage
		} else {
			lang.DefaultLanguage = currentLanguage
		}

		err := example.Page(lang).Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})))

	http.ListenAndServe(":8080", mux)
}
