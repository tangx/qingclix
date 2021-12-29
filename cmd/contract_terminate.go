package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/modules"
)

// contractCmdTerminate 终止合约
var contractCmdTerminate = &cobra.Command{
	Use:   "terminate",
	Short: `终止合约。 如果合约中还有绑定的资源， 则失败。`,
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

	flag.StringVarP(&terminateContracts, "targets", "", "", "合约ID，多个合约ID之间是用 , 进行分割")
	flag.BoolVarP(&terminateConfirm, "confirm", "", false, `确实释放
	必须指定 --confirm=true , 否则不执行。 故意的`)

}

func init() {
	contractCmd.AddCommand(contractCmdTerminate)
}

func terminateContract() {
	if !terminateConfirm {
		logrus.Error("指定 --confirm=true 开启释放开关")
		return
	}

	for _, contract := range targetsFromString(terminateContracts) {
		err := isEmptyContract(contract)
		if err != nil {
			logrus.Errorf("释放合约 (%s) 失败: %v", contract, err)
			continue
		}

		// err = modules.TerminateContract(contract, true)
		// if err != nil {
		// 	logrus.Errorf("释放合约 (%s) 失败: %v", contract, err)
		// 	continue
		// }

		logrus.Infof("释放合约 (%s) 成功", contract)
	}
}

// isEmptyContract 检查合约是否解绑完成所有资源
func isEmptyContract(contract string) error {
	resp := modules.DescContract(contract)
	if resp.RetCode != 0 {
		return fmt.Errorf("%s", resp.Message)
	}

	// 合约分两部分 rc 和 rce
	// 1. 首先查询合约是否存在， 可以获取合约的基本信息， 如区域和类型
	if len(resp.ReservedContractSet) > 0 {
		c := resp.ReservedContractSet[0]
		// output(c)
		// 2. 查询合约内是否绑定有资源
		if len(c.Entries) > 0 {

			for _, entry := range c.Entries {
				if err := isEmptyContractRCE(entry.EntryID); err != nil {
					logrus.Errorf("合约还有资源没有取消绑定: https://console.qingcloud.com/%s/%ss/reserved_contracts/%s/%s", c.ZoneID, c.ResourceType, contract, entry.EntryID)
					return fmt.Errorf("合约还有资源没有取消绑定")
				}
			}
		}
	}

	return nil
}

func isEmptyContractRCE(rce_id string) error {
	resp, err := modules.DescribeReservedResources(rce_id)
	if err != nil {
		return err
	}

	// resp.ReservedResourceSet
	if len(resp.ReservedResourceSet) > 0 {

		return fmt.Errorf("合约绑定资源， 数量 %d", resp.TotalCount)
	}

	return nil
}

func output(v interface{}) {
	b, _ := json.Marshal(v)
	logrus.Infof("output= %s", b)
}
