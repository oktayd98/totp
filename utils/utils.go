package utils

import (
	"os"
	"path/filepath"
)

func GetFilePath() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	dataPath := filepath.Join(homeDir, ".totp", "data.json")

	return dataPath
}
