/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/elsejj/qqwry/qqwry"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the qqwry database",
	Long:  `Update the qqwry database from https://github.com/metowolf/qqwry.dat. you may need to set HTTPS_PROXY environment variable if you are in China.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updating qqwry database to", dbPath, "...")
		qqwry.UpdateMetowolfQQWry(dbPath)
		fmt.Println("Update completed.")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
