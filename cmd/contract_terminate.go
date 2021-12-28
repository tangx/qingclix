package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/modules"
)

// contractCmdTerminate 取消合约
var contractCmdTerminate = &cobra.Command{
	Use:   "terminate",
	Short: "释放合约",
	Run: func(cmd *cobra.Command, args []string) {
		if !isDissocateValidArgs(terminateContracts) {
			_ = cmd.Help()
			return
		}

		// doTerminateContract()

		terminateContract()
	},
}

var (
	terminateContracts string
	terminateConfirm   bool
)

func init() {
	flag := contractCmdTerminate.Flags()

	flag.StringVarP(&terminateContracts, "targets", "", "", "合约ID")
	flag.BoolVarP(&terminateConfirm, "confirm", "", false, "确实释放")

}

func init() {
	contractCmd.AddCommand(contractCmdTerminate)
}

func terminateContract() {
	if !terminateConfirm {
		logrus.Infoln("显示指定 --confirm true 释放资源")
		return
	}

	for _, contract := range targetsFromString(terminateContracts) {
		err := modules.TerminateContract(contract, true)
		if err != nil {
			logrus.Errorf("释放合约失败: %v", err)
		}
	}
}
