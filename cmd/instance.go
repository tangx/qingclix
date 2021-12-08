package cmd

import (
	"github.com/spf13/cobra"
)

// instanceCmd represents the launch command
var instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "购买机器",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(instanceCmd)
}
