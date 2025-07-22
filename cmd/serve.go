/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/elsejj/qqwry/qqwry"
	"github.com/spf13/cobra"
)

var serveListen string
var serveMount string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve as a http server",
	Long: `Serve the qqwry database as a HTTP server.
User can GET the IP information by passing the IP address as a query parameter.
Or POST json/text with the IP address in the body.

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return qqwry.StartHttp(serveListen, serveMount, dbPath)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().StringVarP(&serveListen, "address", "a", "127.0.0.1:11223", "Address to listen on for HTTP requests")
	serveCmd.Flags().StringVarP(&serveMount, "mount", "m", "/", "Mount point for the HTTP server")
}
