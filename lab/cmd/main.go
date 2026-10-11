package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func readLanguageFiles(path string, languages []os.DirEntry) []os.DirEntry {
	files, err := os.ReadDir(path)
	if err != nil {
		return languages
	}

	for _, file := range files {

		if file.IsDir() && file.Name() == "languages" {
			result, err := os.ReadDir(filepath.Join(path, file.Name()))
			if err != nil {
				continue
			}

			languages = append(languages, result...)
			continue
		}

		if file.IsDir() {
			languages = readLanguageFiles(filepath.Join(path, file.Name()), languages)
		}
	}
	return languages
}

func main() {
	files := readLanguageFiles("./", []os.DirEntry{})
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
