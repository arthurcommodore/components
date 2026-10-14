package web

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func returnVersion(filePath string) (error, string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err, ""
	}

	sha := sha256.Sum256(content)
	return nil, hex.EncodeToString(sha[:6])
}

func returnContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func assetURL(filePath string) string {
	err, version := returnVersion(filePath)
	if err != nil {
		version = "none"
	}
	return filePath + "?v=" + version
}
