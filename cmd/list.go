/*
Copyright © 2023 Oktay Dönmez <oktaydonmez98@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/oktayd98/totp/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all otp records.",
	Long:  `Lists all names of saved otp records.`,
	Run: func(cmd *cobra.Command, args []string) {
		showKeys, _ := cmd.Flags().GetBool("show-keys")
		showOTPs, _ := cmd.Flags().GetBool("show-otps")

		list(showKeys, showOTPs)
	},
}

func list(showKeys bool, showOTPs bool) {
	filePath := utils.GetFilePath()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File could not find. You must run create command first.")
		return
	}

	data := utils.ReadJSONFromFile(filePath)

	tableHeader := table.Row{"#", "Name"}
	tableRows := []table.Row{}

	for i, v := range data.OTPs {
		tableRows = append(tableRows, table.Row{i + 1, v.Name})
	}

	if showOTPs {
		tableHeader = append(tableHeader, "OTP")
		for i, v := range data.OTPs {
			otp := utils.GenerateOTP(v.Key)
			tableRows[i] = append(tableRows[i], otp)
		}
	}

	if showKeys {
		tableHeader = append(tableHeader, "Key")
		for i, v := range data.OTPs {
			tableRows[i] = append(tableRows[i], v.Key)
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(tableHeader)
	t.AppendRows(tableRows)
	t.Render()
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("show-keys", "k", false, "Show secret keys")
	listCmd.Flags().BoolP("show-otps", "o", true, "Show OTPs")
}
