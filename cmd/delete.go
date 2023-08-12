/*
Copyright © 2023 Oktay Dönmez <oktaydonmez98@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/oktayd98/totp/types"
	"github.com/oktayd98/totp/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an otp record",
	Long:  `Deletes an otp record by name.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		delete(name)
	},
}

func delete(name string) {
	filePath := utils.GetFilePath()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File could not find. You must run create command first.")
		return
	}

	existingData := utils.ReadJSONFromFile(filePath)
	var newData types.OTPData

	for _, v := range existingData.OTPs {
		if v.Name != name {
			newData.OTPs = append(newData.OTPs, v)
		}
	}

	fmt.Println(newData)
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("name", "n", "", "Name of OTP")

	deleteCmd.MarkFlagRequired("name")
}
