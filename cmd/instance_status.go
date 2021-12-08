package cmd

import (
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
	instanceCmdStatus.Flags().StringVarP(&switchStatusTargets, "targets", "t", "", "开关机对象")
	instanceCmdStatus.Flags().StringVarP(&switchStatusAction, "action", "", "", "开关机行（显示指定）: start | stop")
	instanceCmdStatus.Flags().BoolVarP(&stopForce, "force", "", false, "强制关机")
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

	modules.StopInstance(ins, force)
}

func switchInstanceStatus() {
	ins := targetsFromString(switchStatusTargets)

	switch switchStatusAction {
	case "stop":
		stopInstances(ins...)
	}
}

func isValidSwitchStatus() bool {
	if len(switchStatusTargets) == 0 {
		return false
	}

	return true
}
