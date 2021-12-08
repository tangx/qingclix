package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/modules"
)

// contractCmdDissocite represents the contract command
var contractCmdDissocite = &cobra.Command{
	Use:   "dissocite",
	Short: "为资源解绑合约",
	Run: func(cmd *cobra.Command, args []string) {
		if !isDissocateValidArgs() {
			_ = cmd.Help()
			return
		}

		doDissociteContract()
	},
}

var (
	dissociteTargetType string
	dissociteTargets    string
	fileToDissocite     string
)

func init() {
	contractCmd.AddCommand(contractCmdDissocite)
	contractCmdDissocite.Flags().StringVarP(&dissociteTargets, "target", "t", "", "为指定资源解绑合约, i-xxx,i-yyy,i-zzz")
	contractCmdDissocite.Flags().StringVarP(&dissociteTargets, "file", "f", "", `为指定资源解绑合约, 通过文件导入。 /path/to/file
文件内可以多行， 每行可以多个资源， 资源之间以 , 分隔。`)
	contractCmdDissocite.Flags().StringVarP(&dissociteTargetType, "type", "", "instance", "为指定资源制定类型，所有资源必须为相同类型。 可选项: instance")
}

func isDissocateValidArgs() bool {
	if len(dissociteTargets) == 0 && len(fileToDissocite) == 0 {
		return false
	}

	return true
}

func targetsFromString(str string) []string {
	return strings.Split(dissociteTargets, ",")
}

func targetsFromFile(file string) (targets []string) {

	f, err := os.Open(file)
	if err != nil {
		logrus.Errorf("open file %s failed: %v", err)
		return
	}
	defer f.Close()

	buf := bufio.NewScanner(f)
	for buf.Scan() {
		line := buf.Text()

		ret := targetsFromString(line)
		targets = append(targets, ret...)
	}

	return
}

func doDissociteContract() {

	instances := targetsFromFile(fileToDissocite)
	instances = append(instances, targetsFromString(dissociteTargets)...)

	for _, instance := range instances {
		instance = strings.TrimSpace(instance)
		if len(instance) == 0 {
			continue
		}

		contract := getResourceContract(instance)
		if contract == "" {
			logrus.Warnf("资源 %s 没有绑定合约", instance)
			continue
		}

		ok := modules.DissociateContract(contract, instance)
		if !ok {
			logrus.Errorf("解绑资源 %s 合约 %s 失败", instance, contract)
			continue
		}

		logrus.Infof("解绑资源 %s 合约 %s 成功", instance, contract)
	}
}

func getResourceContract(res string) (contract string) {

	switch dissociteTargetType {
	case "instance":
		return getInstanceContract(res)
	}

	return
}

func getInstanceContract(instance string) (contract string) {

	resp := modules.DescInstance(instance)

	for _, ins := range resp.InstanceSet {
		return ins.ReservedContract
	}

	return
}