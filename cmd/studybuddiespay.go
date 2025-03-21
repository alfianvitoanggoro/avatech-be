package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var studdybuddiespayCmd = &cobra.Command{
	Use:   "studybuddiespay",
	Short: "studybuddiespay",
	Long:  `studybuddiespay`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("studybuddiespay")
	},
}

func init() {
	rootCmd.AddCommand(studdybuddiespayCmd)
}
