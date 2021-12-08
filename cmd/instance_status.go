package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/modules"
)

// instanceCmd represents the launch command
var instanceCmdStatus = &cobra.Command{
	Use:   "status",
	Short: "关机",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !isValidSwitchStatus() {
			_ = cmd.Help()
		}
		switchInstanceStatus()
	},
}

func init() {
	instanceCmd.AddCommand(instanceCmdStatus)
	instanceCmdStatus.Flags().StringVarP(&switchStatusTargets, "targets", "t", "", "开关机对象， 以逗号分隔: ins-123,ins-yyy,ins-zzz")
	instanceCmdStatus.Flags().StringVarP(&switchStatusAction, "action", "", "", "开关机行（显示指定）: start | stop | restart")
	instanceCmdStatus.Flags().BoolVarP(&stopForce, "force", "", false, "强制关机, 仅 action=stop 有效")
}

var (
	switchStatusTargets string
	switchStatusAction  string
	stopForce           bool
)

func stopInstances(ins ...string) {
	force := 0
	if stopForce {
		force = 1
	}

	id := modules.StopInstance(ins, force)
	logrus.Infof("关机任务已提交， 任务id: %s", id)
}

func restartInstances(ins []string, action modules.InstanceAction) {

	id := modules.StartOrRestartInstances(ins, action)
	logrus.Infof("关机任务已提交， 任务id: %s", id)
}
func switchInstanceStatus() {
	ins := targetsFromString(switchStatusTargets)

	action := modules.InstanceAction(switchStatusAction)

	switch action {
	case modules.InstanceActionStop:
		stopInstances(ins...)
	case modules.InstanceActionStart, modules.InstanceActionRestart:
		restartInstances(ins, action)
	}

}

func isValidSwitchStatus() bool {
	if len(switchStatusTargets) == 0 {
		return false
	}

	return true
}
