/*
Copyright © 2023 Oktay Dönmez <oktaydonmez98@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "totp",
	Short: "Totp is a CLI application that helps you to manage your OTPs.",
	Long: `Totp is a CLI application that helps you to manage your OTPs.
You can manage easily your one time passwords with Totp from command line.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
