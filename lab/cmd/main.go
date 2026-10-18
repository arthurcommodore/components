package main

import (
	"fmt"

	"github.com/arthurcommodore/components/language"
)

func main() {
	language := language.NewLanguage("lab/languages", "pt")

	tradPt := language.Translator("pt")
	text := tradPt("teste")
	fmt.Println(text)
}
