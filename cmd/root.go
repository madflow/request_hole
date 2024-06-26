// Package cmd contains all of the CLI commands
package cmd

import (
	"net/http"

	"github.com/spf13/cobra"
)

var (
	Address      string
	BuildInfo    map[string]string
	Details      bool
	LogFile      string
	Port         int
	ResponseCode int
	Web          bool
	WebAddress   string
	WebPort      int
	StaticFS     http.FileSystem
)

var rootCmd = &cobra.Command{
	Use:   "rh",
	Short: "A CLI for an ephemeral API endpoint",
	Long: `rh: Request Hole
This CLI tool will let you create a temporary API endpoint for testing purposes.`,
}

func Execute(buildInfo map[string]string, staticFS http.FileSystem) error {
	BuildInfo = buildInfo
	StaticFS = staticFS
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8080, "sets the port for the endpoint")
	rootCmd.PersistentFlags().StringVarP(&Address, "address", "a", "localhost", "sets the address for the endpoint")
	rootCmd.PersistentFlags().IntVarP(&ResponseCode, "response_code", "", 200, "sets the response code")
	rootCmd.PersistentFlags().BoolVar(&Details, "details", false, "shows header details in the request")
	rootCmd.PersistentFlags().StringVar(&LogFile, "log", "", "writes incoming requests to the specified log file (example: --log rh.log)")

	// Web server renderer
	rootCmd.PersistentFlags().BoolVar(&Web, "web", false, "runs the web UI to show incoming requests")
	rootCmd.PersistentFlags().StringVar(&WebAddress, "web_address", "localhost", "sets the address for the web UI")
	rootCmd.PersistentFlags().IntVar(&WebPort, "web_port", 8081, "sets the port for the web UI")
}
