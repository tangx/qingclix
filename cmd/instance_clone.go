package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/types"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "clone 目标机器",
	Long: `指定一台已运行的服务器，购买一台相同配置的服务器。
注意: 由于青云无法指定磁盘挂载到服务器上的盘符信息。因此磁盘挂载顺序不可预知。`,
	Run: func(cmd *cobra.Command, args []string) {
		instanceCloneMain()
	},
}

var (
	instance_clone_target string
)

func init() {
	instanceCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().StringVarP(&instance_clone_target, "instance", "i", "", "指定克隆的源主机")

}

func instanceCloneMain() {
	cli := types.Client{}

	if instance_clone_target == "" {
		logrus.Errorln("没有指定克隆的目标机")
	}

	item, err := cloneConfig(cli, instance_clone_target)
	if err != nil {
		panic(err)
	}

	InstallInstance(item)
}
