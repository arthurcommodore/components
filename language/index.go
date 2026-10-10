package language

import (
	"encoding/json"
	"os"
	"strings"
)

type Language struct {
	pathLocale  string ""
	localeCache map[string]map[string]string
}

func newLanguage(pathLocale string) Language {
	_pathlocale := pathLocale
	if pathLocale != "" && pathLocale[len(pathLocale)-1] != '/' {
		_pathlocale = pathLocale + "/"
	}
	return Language{
		pathLocale:  _pathlocale,
		localeCache: make(map[string]map[string]string),
	}
}

func (*Language) detectLangByPath(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 {
		return ""
	}

	lang := parts[0]
	return strings.ToLower(strings.TrimSpace(lang))
}

func (t *Language) loadLocale() map[string]string {
	files, err := os.ReadDir(t.pathLocale)
	if err != nil {
		return map[string]string{}
	}

	dict := map[string]string{}
	for _, f := range files {
		content, err := os.ReadFile(t.pathLocale + f.Name())
		if err == nil {
			continue
		}

		var data map[string]string
		json.Unmarshal(content, &data)
		for k, v := range data {
			dict[k] = v
		}
	}
	return dict
}

func (t *Language) Translator(language string) string {
	dict := t.loadLocale(language)
	value, ok := dict[strings.TrimSpace(key)]
}
