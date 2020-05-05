package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/modules"
	"github.com/tangx/qingyun-sdk-go/qingyun"
)

// instanceCmd represents the launch command
var instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "购买机器",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// instanceCmd represents the launch command
var instanceCmdrun = &cobra.Command{
	Use:   "run",
	Short: "购买机器",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		LaunchMain()
	},
}

func init() {
	rootCmd.AddCommand(instanceCmd)
	instanceCmd.AddCommand(instanceCmdrun)
}

func LaunchMain() {
	item := chooseItem()

	// run instance
	// instParams := (*qingyun.RunInstancesRequest)(unsafe.Pointer(&item.Instance))

	instParams := qingyun.RunInstancesRequest{
		ImageID:       item.Instance.ImageID,
		InstanceName:  item.Instance.InstanceName,
		InstanceClass: item.Instance.InstanceClass,
		CPU:           item.Instance.CPU,
		Memory:        item.Instance.Memory,
		LoginMode:     item.Instance.LoginMode,
		LoginKeypair:  item.Instance.LoginKeypair,
		Zone:          item.Instance.Zone,
		Vxnets:        item.Instance.Vxnets,
	}
	instID := modules.RunInstance(instParams)
	logrus.Debugf("New instance id = %s", instID)

	if instID == "" {
		logrus.Fatalf("Error: Create Instance Failed")
	}

	// create volumes
	vols := []string{}
	for i, v := range item.Volumes {
		params := qingyun.CreateVolumesRequest{
			VolumeName: fmt.Sprintf("%s-%d", instParams.InstanceName, i),
			VolumeType: v.VolumeType,
			Size:       v.Size,
			Zone:       item.Instance.Zone,
		}
		vol, _ := modules.CreateVolume(params)
		vols = append(vols, vol)
	}

	// attach
	for _, vol := range vols {
		modules.AttachVolume(instID, vol)
	}
}
