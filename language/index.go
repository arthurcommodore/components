package language

import (
	"encoding/json"
	"os"
	"strings"
)

type Language struct {
	pathLocale      string ""
	localeCache     map[string]map[string]string
	defaultLanguage string
}

func NewLanguage(pathLocale string, defaultLanguage string) Language {
	_pathlocale := pathLocale
	if pathLocale != "" && pathLocale[len(pathLocale)-1] != '/' {
		_pathlocale = pathLocale + "/"
	}
	language := Language{
		pathLocale:      _pathlocale,
		localeCache:     make(map[string]map[string]string),
		defaultLanguage: defaultLanguage,
	}
	language.injectLocaleCache()
	return language
}

func (*Language) detectLangByPath(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 {
		return ""
	}

	lang := parts[0]
	return strings.ToLower(strings.TrimSpace(lang))
}

func (t *Language) loadLocale(path string) map[string]string {
	files, err := os.ReadDir(path)
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

func (l *Language) injectLocaleCache() {
	langsDir, error := os.ReadDir(l.pathLocale)
	if error != nil {

	}
	for _, dir := range langsDir {
		dirs, error := os.ReadDir(dir.Name())
		if error != nil {

		}
		for _, dir := range dirs {
			l.localeCache[dir.Name()] = l.loadLocale(dir.Name())
		}
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
