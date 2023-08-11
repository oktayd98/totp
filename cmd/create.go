/*
Copyright © 2023 Oktay Dönmez <oktaydonmez98@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/oktayd98/totp/types"
	"github.com/oktayd98/totp/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates new otp record",
	Long:  `Creates a new otp records with given names.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		name, _ := cmd.Flags().GetString("name")

		create(key, name)
	},
}

func create(key string, name string) {
	filePath := utils.GetFilePath()
	dirPath := filepath.Dir(filePath)

	if err := os.MkdirAll(dirPath, 0700); err != nil {
		panic(err)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		initialData := types.OTPData{OTPs: []types.OTP{}}
		saveJSONToFile(initialData, filePath)
	}

	existingData := readJSONFromFile(filePath)

	newOTP := types.OTP{
		Key:       key,
		Name:      name,
		CreatedAt: time.Now().Unix(),
	}

	existingData.OTPs = append(existingData.OTPs, newOTP)

	saveJSONToFile(existingData, filePath)
}

func saveJSONToFile(data types.OTPData, filePath string) {
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

func readJSONFromFile(filePath string) types.OTPData {
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

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("key", "k", "", "Secret key of OTP.")
	createCmd.Flags().StringP("name", "n", "", "Name of OTP.")

	createCmd.MarkFlagRequired("key")
	createCmd.MarkFlagRequired("name")
}
