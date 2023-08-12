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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets otp by name",
	Long:  `Gets otp by name.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		get(name)
	},
}

func get(name string) {
	filePath := utils.GetFilePath()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File could not find. You must run create command first.")
		return
	}

	data := utils.ReadJSONFromFile(filePath)
	var record types.OTP

	for _, v := range data.OTPs {
		if v.Name == name {
			record = v
			break
		}
	}

	otpString := utils.GenerateOTP(record.Key)

	fmt.Println(otpString)
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("name", "n", "", "Name of OTP.")

	getCmd.MarkFlagRequired("name")
}
