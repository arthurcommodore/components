package language

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Language struct {
	pathLocale      string ""
	localeCache     map[string]map[string]string
	defaultLanguage string
}

func NewLanguage(pathLocale string, defaultLanguage string) Language {
	language := Language{
		pathLocale:      pathLocale,
		localeCache:     make(map[string]map[string]string),
		defaultLanguage: defaultLanguage,
	}
	language.injectLocaleCache()
	return language
}

func (l *Language) loadLocale(path string) map[string]string {
	files, err := os.ReadDir(path)
	if err != nil {
		return map[string]string{}
	}

	dict := map[string]string{}
	for _, f := range files {
		content, err := os.ReadFile(filepath.Join(path, f.Name()))
		if err != nil {
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

func (l *Language) injectLocaleCache() {
	langsDir, error := os.ReadDir(l.pathLocale)
	if error != nil {
		log.Fatalf("Error injectLocaleCache language pathLocale: %s", l.pathLocale)
	}
	for _, lang := range langsDir {
		if !lang.IsDir() {
			continue
		}
		l.localeCache[lang.Name()] = l.loadLocale(filepath.Join(l.pathLocale, lang.Name()))
	}
}

func (l *Language) Translator(lang string) func(key string, params ...string) string {
	translatorLang := l.localeCache[lang]
	return func(key string, params ...string) string {
		value, ok := translatorLang[strings.TrimSpace(key)]
		if !ok {
			return "[" + key + "]"
		}

		for i := 0; i < len(params)-1; i += 2 {
			value = strings.ReplaceAll(value, params[i], params[i+1])
		}
		return value
	}
}
