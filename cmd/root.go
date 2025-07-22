/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

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
	dbPath = findDB()
	if dbPath == "" {
		println("qqwry.dat not found, please specify the path with --db option or place it in a standard location.")
		os.Exit(1)
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
	rootCmd.PersistentFlags().StringVarP(&dbPath, "db", "d", "qqwry.dat", "Path to the qqwry.dat database file")
}

func findDB() string {
	if isFile(dbPath) {
		return dbPath
	}
	findPlaces := []string{
		"/usr/local/share/qqwry.dat",
		"/usr/share/qqwry.dat",
		"/usr/local/qqwry.dat",
		"/usr/qqwry.dat",
		"/etc/qqwry.dat",
		"/qqwry.dat",
		"qqwry.dat",
		"$HOME/qqwry.dat",
		"$HOME/.local/share/qqwry.dat",
		"$HOME/.qqwry.dat",
		"$HOME/.qqwry/qqwry.dat",
		"$HOME/.config/qqwry.dat",
		"$HOME/.config/qqwry/qqwry.dat",
		"$HOME/AppData/Local/qqwry.dat",
		"$HOME/AppData/Roaming/qqwry.dat",
		"$HOME/AppData/Local/qqwry/qqwry.dat",
		"$HOME/AppData/Roaming/qqwry/qqwry.dat",
	}
	for _, place := range findPlaces {
		place = os.ExpandEnv(place)
		if isFile(place) {
			return place
		}
	}
	return ""
}
