package cmd

import (
	"fmt"

	"github.com/alfianvitoanggoro/avatech/libs/avapay"
	"github.com/spf13/cobra"
)

var avapayCmd = &cobra.Command{
	Use:   "avapay",
	Short: "avapay",
	Long:  `avapay`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		default:
			cmd.Help()
		case avapay.Name != "":
			avapay.PayerName(avapay.Name)
		case avapay.PayStatus:
			fmt.Println("Status: ", avapay.PayStatus)
		}
	},
}

func init() {
	rootCmd.AddCommand(avapayCmd)

	avapayCmd.Flags().StringVarP(&avapay.Name, "name", "n", "", "Execute for get name")

	avapayCmd.Flags().BoolVarP(&avapay.PayStatus, "status", "s", false, "Execute for get pay status")
}
