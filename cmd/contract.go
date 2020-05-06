package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// contractCmd represents the contract command
var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("contract called")
	},
}

// contractCmdApplyfor represents the contract command
var contractCmdApplyfor = &cobra.Command{
	Use:   "apply-for",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApplyForResource()
	},
}

func init() {
	rootCmd.AddCommand(contractCmd)
	contractCmdApplyfor.Flags().StringVarP(&target_resource, "target", "t", "", "为指定资源购买合约")
}

var (
	target_resource string
)

func ApplyForResource() {
	if target_resource == "" {
		logrus.Infof("No Target to Apply Contract for")
		return
	}

	applyForResource(target_resource)

}

func applyForResource(resource string) {
}
