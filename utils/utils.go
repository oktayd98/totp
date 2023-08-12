package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/oktayd98/totp/types"
)

func GetFilePath() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	dataPath := filepath.Join(homeDir, ".totp", "data.json")

	return dataPath
}

func SaveJSONToFile(data types.OTPData, filePath string) {
	file, err := os.Create(filePath)

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		panic(err)
	}
}

func ReadJSONFromFile(filePath string) types.OTPData {
	data := types.OTPData{}

	fi, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(fi).Decode(&data)

	if err != nil {
		panic(err)
	}

	return data
}
