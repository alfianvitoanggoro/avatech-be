package debug

import (
	"github.com/spf13/cobra"
)

var avapayCmd = &cobra.Command{
	Use:   "avapay",
	Short: "Debug Avapay",
	Long:  `Debug Avapay.`,
	Run: func(cmd *cobra.Command, args []string) {
		message := "Debug Avapay Disburse"
		success("Message : ", message)
	},
}

func init() {
	RootCmd.AddCommand(avapayCmd)
}
