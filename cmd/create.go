/*
Copyright © 2023 Oktay Dönmez <oktaydonmez98@gmail.com>
*/
package cmd

import (
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
		utils.SaveJSONToFile(initialData, filePath)
	}

	existingData := utils.ReadJSONFromFile(filePath)

	newOTP := types.OTP{
		Key:       key,
		Name:      name,
		CreatedAt: time.Now().Unix(),
	}

	existingData.OTPs = append(existingData.OTPs, newOTP)

	utils.SaveJSONToFile(existingData, filePath)
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("key", "k", "", "Secret key of OTP.")
	createCmd.Flags().StringP("name", "n", "", "Name of OTP.")

	createCmd.MarkFlagRequired("key")
	createCmd.MarkFlagRequired("name")
}
