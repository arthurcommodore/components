package main

import (
	"fmt"

	"github.com/arthurcommodore/components/language"
)

func main() {
	l := language.NewLanguage("lab/languages", "pt")

	text := l.Translate("teste")
	fmt.Println(text)
}
