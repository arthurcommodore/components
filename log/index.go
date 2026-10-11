package log

import (
	"log"
	"os"
	"path/filepath"
)

var PathFile = "./"
var LogFile *os.File

func init() {
	if err := SetPathFile(PathFile); err != nil {
		log.Println("Erro ao abrir arquivo de log:", err)
	}
}

func Close() error {
	if LogFile == nil {
		return nil
	}

	return LogFile.Close()
}

func SetPathFile(pathFile string) error {
	if err := os.MkdirAll(pathFile, 0755); err != nil {
		return err
	}
	filePath := filepath.Join(pathFile, "server.log")

	newFile, err := os.OpenFile(
		filePath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	if LogFile != nil {
		_ = LogFile.Close()
	}

	LogFile = newFile
	PathFile = pathFile
	log.SetOutput(newFile)

	return nil
}

func LogSystem(message string, args ...any) {
	log.Printf(message, args...)
}
