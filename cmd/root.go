package cmd

import (
	"os"

	"github.com/alfianvitoanggoro/avatech/cmd/debug"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "avatech",
	Short: "Avatech BE",
	Long:  `Avatech BE is a backend of apps avatech platform`,
}

func init() {
	rootCmd.AddCommand(debug.RootCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
