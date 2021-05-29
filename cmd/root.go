// Package cmd contains all of the CLI commands
package cmd

import "github.com/spf13/cobra"

var (
	Address      string
	Port         int
	ResponseCode int
)

var rootCmd = &cobra.Command{
	Use:   "rq",
	Short: "A CLI for an ephemeral API endpoint",
	Long: `rq: Request Hole
This CLI tool will let you create a temporary API endpoint for testing purposes.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8080, "sets the port for the endpoint")
	rootCmd.PersistentFlags().StringVarP(&Address, "address", "a", "localhost", "sets the address for the endpoint")
	rootCmd.PersistentFlags().IntVarP(&ResponseCode, "response_code", "r", 200, "sets the response code")
}