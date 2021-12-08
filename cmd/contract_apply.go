package cmd

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/modules"
	"github.com/tangx/qingclix/sdk-go/qingyun"
)

// contractCmdApplyfor represents the contract command
var contractCmdApplyfor = &cobra.Command{
	Use:   "apply-for",
	Short: "对已有资源创建合约",
	Run: func(cmd *cobra.Command, args []string) {
		// _ = cmd.Help()
		if len(target_resource) == 0 {
			_ = cmd.Help()
			return
		}

		ApplyForResource()
	},
}

func init() {
	contractCmd.AddCommand(contractCmdApplyfor)
	contractCmdApplyfor.Flags().StringVarP(&target_resource, "target", "t", "", "为指定资源购买合约, i-xxx,i-yyy,i-zzz")
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

// applyForResource contract
// resources i-xxx,i-yyy,i-zzz
func applyForResource(resources string) {

	for _, res := range strings.Split(resources, ",") {
		logrus.Infof("apply concat for %s ... ", res)

		params := qingyun.ApplyReservedContractWithResourcesRequest{
			Resources: []string{res},
			Months:    1,
			AutoRenew: 1,
		}

		contractid := modules.ApplyContractWithResources(params)

		if len(contractid) != 0 {
			fmt.Printf("(contractid) %s", contractid)
		}
		fmt.Println("")
	}

}
