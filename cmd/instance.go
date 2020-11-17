package cmd

import (
	"fmt"
	"sync"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/cmd/configure"
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
		_ = cmd.Help()
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

var (
	instance_count int
	start_with     int
)

func init() {
	rootCmd.AddCommand(instanceCmd)
	instanceCmd.AddCommand(instanceCmdrun)
	instanceCmdrun.Flags().IntVarP(&instance_count, "count", "c", 1, "新购实例数量")
	instanceCmdrun.Flags().IntVarP(&start_with, "start-with", "", 0, "服务器名称起始编号")

}

var wg sync.WaitGroup

func LaunchMain() {

	item := configure.ChooseItem()

	instName := item.Instance.InstanceName
	for i := 0; i < instance_count; i++ {
		wg.Add(1)

		if i == 0 {
			item.Instance.InstanceName = instName + fmt.Sprintf("-%d", start_with)
		} else {
			item.Instance.InstanceName = instName + fmt.Sprintf("-%d", i+start_with)
		}

		go RunInstance(item)
	}
	wg.Wait()
}

func RunInstance(item configure.ClixItem) {
	defer wg.Done()

	logrus.WithFields(logrus.Fields{
		"ClixJobID": uuid.NewV4().String(),
	}).Infof("Create Instance :%s", item.Instance.InstanceName)

	// struct 强制类型转换 ， 此处不可用 :
	// https://www.cnblogs.com/hitfire/articles/6363696.html
	// instParams := (*qingyun.RunInstancesRequest)(unsafe.Pointer(&item.Instance))

	instParams := qingyun.RunInstancesRequest{
		ImageID:       item.Instance.ImageID,
		InstanceName:  item.Instance.InstanceName,
		InstanceClass: item.Instance.InstanceClass,
		CPU:           item.Instance.CPU,
		CPUModel:      item.Instance.CPUModel,
		Memory:        item.Instance.Memory,
		LoginMode:     item.Instance.LoginMode,
		LoginKeypair:  item.Instance.LoginKeypair,
		Zone:          item.Instance.Zone,
		Vxnets:        item.Instance.Vxnets,
		OsDiskSize:    item.Instance.OsDiskSize,
		GPU:           item.Instance.GPU,
		GPUClass:      item.Instance.GPUClass,
	}
	instID := modules.RunInstance(instParams)
	logrus.Debugf("New instance id = %s", instID)

	if instID == "" {
		logrus.Fatalf("Error: Create Instance Failed")
	}

	// create volumes
	vols := []string{}
	for i, v := range item.Volumes {

		volname := instParams.InstanceName
		if i != 0 {
			volname = fmt.Sprintf("%s-%d", instParams.InstanceName, i)
		}

		params := qingyun.CreateVolumesRequest{
			VolumeName: volname,
			VolumeType: v.VolumeType,
			Size:       v.Size,
			Zone:       item.Instance.Zone,
		}
		vol, _ := modules.CreateVolume(params)
		vols = append(vols, vol)
	}

	// attach
	for _, vol := range vols {
		_ = modules.AttachVolume(instID, vol)
	}

	// contract apply,lease and associate
	modules.CreateContract(instID, item.Contract.AutoRenew, item.Contract.Months, item.Instance.Zone)
	for _, vol := range vols {
		modules.CreateContract(vol, item.Contract.AutoRenew, item.Contract.Months, item.Instance.Zone)
	}
}
