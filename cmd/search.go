/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/elsejj/qqwry/qqwry"
	"github.com/spf13/cobra"
)

var searchOutput string
var searchOutputFormat string
var showPerformance bool

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search IP addresses in the qqwry database",
	Long: `Search for IP addresses in the qqwry database. 
You can pass IP addresses as arguments, which will be looked up in the database.
Or you can pass files containing IP addresses, line containing IP addresses will be replaced with the search result.
When no arguments are provided, it will read from standard input.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		loadT1 := time.Now()
		db, err := qqwry.NewDb(dbPath)
		if err != nil {
			cmd.Println("Error initializing database:", err)
			return
		}
		loadT2 := time.Now()
		if showPerformance {
			fmt.Fprintln(os.Stderr, "Database loaded in", loadT2.Sub(loadT1).Microseconds(), "us")
		}

		searchOutputFile := os.Stdout
		if len(searchOutput) > 0 {
			fp, err := os.Create(searchOutput)
			if err != nil {
				cmd.Println("Error creating output file:", err)
				return
			}
			searchOutputFile = fp
		}
		defer searchOutputFile.Close()

		formatter := currentFormatter()

		if len(args) == 0 {
			qqwry.SearchReplace(db, os.Stdin, searchOutputFile, formatter)
		} else {
			for _, arg := range args {
				if isFile(arg) {
					file, err := os.Open(arg)
					if err != nil {
						cmd.Println("Error opening file:", arg, " error:", err)
						continue
					}
					defer file.Close()
					qqwry.SearchReplace(db, file, searchOutputFile, formatter)
				} else if qqwry.IsIpV4(arg) {
					searchT1 := time.Now()
					result := qqwry.SearchIp(db, arg)
					searchT2 := time.Now()
					if showPerformance {
						fmt.Fprintln(os.Stderr, "Search for", arg, "took", searchT2.Sub(searchT1).Nanoseconds(), "ns")
					}
					searchOutputFile.WriteString(formatter(result) + "\n")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchCmd.Flags().StringVarP(&searchOutput, "output", "o", "", "Output file for search results, if not set, will print to standard output")
	searchCmd.Flags().StringVarP(&searchOutputFormat, "format", "f", "csv", "Output format for search results (e.g., json, csv)")
	searchCmd.Flags().BoolVarP(&showPerformance, "time", "t", false, "Show performance metrics for the search operation")

}

func currentFormatter() func(result qqwry.SearchResult) string {
	switch searchOutputFormat {
	case "json":
		return qqwry.FormatJSON
	default:
		return qqwry.FormatCSV
	}
}
