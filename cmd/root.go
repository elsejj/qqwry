/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/elsejj/qqwry/qqwry"
	"github.com/spf13/cobra"
)

var dbPath string

func isFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qqwry",
	Short: "A IP lookup tool based on qqwry.dat",
	Long:  `qqwry is a command line tool for looking up IP addresses using the qqwry.dat database.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if !isFile(dbPath) {
		// if the file does not exist, print an error message and exit
		fmt.Println("Error: The database file does not exist:", dbPath)
		fmt.Println("Download the database file from https://github.com/metowolf/qqwry.dat [y/n] ?")
		var response string
		fmt.Scanln(&response)
		if response == "y" || response == "Y" {
			fmt.Println("Downloading the database file...")
			qqwry.UpdateMetowolfQQWry(dbPath)
		}
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.qqwry.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&dbPath, "db", "d", defaultDBPath(), "Path to the qqwry.dat database file")
}

func defaultDBPath() string {
	// 1. check if the QQWRY_DAT environment variable is set
	if envPath := os.Getenv("QQWRY_DAT"); envPath != "" {
		return envPath
	}
	// 2. return the os-specific default path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "qqwry.dat" // fallback to current directory if home dir cannot be determined
	}
	configDir := ""
	switch runtime.GOOS {
	case "windows":
		configDir = path.Join(homeDir, "AppData", "Local", "qqwry")
	default:
		configDir = path.Join(homeDir, ".config", "qqwry")
	}

	os.MkdirAll(configDir, 0755) // ensure the directory exists
	qqwryDat := path.Join(configDir, "qqwry.dat")
	return qqwryDat
}
