package cmd

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/types"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "对预设服务器配置进行管理",
	Long: `配置管理
clone: clone 指定服务器的配置
interactive: 交互式创建配置
delete: 删除配置
`,
	Run: func(cmd *cobra.Command, args []string) {
		configureMain()
	},
}

var (
	configure_clone_target string
	configure_interactive  bool
	configure_label        string
	configure_initial      bool
)

func init() {
	instanceCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringVarP(&configure_clone_target, "clone", "", "", "克隆目标服务器配置")
	configureCmd.Flags().StringVarP(&configure_label, "label", "l", "", "使用自定义label替代默认生成规则")
	configureCmd.Flags().BoolVarP(&configure_interactive, "interactive", "i", false, "交互问答生成配置")
	configureCmd.Flags().BoolVarP(&configure_initial, "initial", "", false, "初始化配置文件（注意: 覆盖现有文件)")

}

func configureMain() {
	if configure_initial {
		InitialConfig()
		return
	}

	if configure_interactive {

		return
	}

	if configure_clone_target != "" {
		configureClone(configure_clone_target)
		return
	}
}

func appendItemConfig(name string, item ItemConfig) {
	presetConfig := LoadPresetConfig()
	presetConfig.Configs[name] = item

	SaveConfigToFile(presetConfig)
}

func configureClone(instance string) {
	cli := types.Client{}

	if instance == "" {
		logrus.Errorln("没有指定克隆的目标机")
	}

	item, err := cloneConfig(cli, instance)
	if err != nil {
		logrus.Errorln(err)
	}

	if configure_label == "" {
		configure_label = item.Instance.InstanceName + "_clone"
	}
	appendItemConfig(configure_label, item)

}

func cloneConfig(cli types.Client, instance string) (config ItemConfig, err error) {

	insParams, volumes, contract := cloneInstance(cli, instance)

	// clone instance
	config.Instance = insParams

	// clone volumes
	volumesParams := cloneVolume(cli, volumes)
	config.Volumes = volumesParams

	// clone contract
	contractParams := cloneContract(cli, contract)
	config.Contract = contractParams

	return
}

func cloneInstance(cli types.Client, instance string) (params types.RunInstancesRequest, volumes []string, contract string) {
	descParams := types.DescribeInstancesRequest{
		Instances: []string{instance},
		Verbose:   1,
	}
	descResp, err := cli.DescribeInstances(descParams)
	if err != nil {
		logrus.Errorf("%s", err)
		return
	}
	if descResp.TotalCount == 0 {
		msg := fmt.Sprintf("找不到指定主机: %s", instance)
		logrus.Infof("%s", errors.New(msg))
		return
	}

	// 遍历加入的网络
	target := descResp.InstanceSet[0]
	vxnets := []string{}
	for _, vxnet := range target.Vxnets {
		vxnets = append(vxnets, vxnet.VxnetID)
	}

	var keypair string

	if len(target.KeypairIDS) == 0 {
		keypair = ""
	} else {
		keypair = target.KeypairIDS[0]
	}
	logrus.Debugf("Keypair 数量: %d, keypair %s", len(target.KeypairIDS), target.KeypairIDS)

	// 初始化服务器请求变量
	params = types.RunInstancesRequest{
		ImageID:      target.Image.ImageID,
		InstanceType: target.InstanceType,
		InstanceName: target.InstanceName + "_clone",
		Zone:         target.ZoneID,
		Vxnets:       vxnets,
		LoginKeypair: keypair,
		LoginMode:    "keypair",
	}

	volumes = target.VolumeIDS
	contract = target.ReservedContract

	return
}

func cloneVolume(cli types.Client, volumes []string) (params []types.CreateVolumesRequest) {
	descParams := types.DescribeVolumesRequest{
		Volumes: volumes,
	}
	resp, err := cli.DescribeVolumes(descParams)
	if err != nil {
		logrus.Errorln(err)
	}

	for _, volume := range resp.DescribeVolumeSet {
		volParams := types.CreateVolumesRequest{
			Size:       volume.Size,
			VolumeType: volume.VolumeType,
		}
		params = append(params, volParams)
	}

	return
}

func cloneContract(cli types.Client, contract string) (params types.ApplyReservedContractWithResourcesRequest) {
	descParams := types.DescribeReservedContractsRequest{
		ReservedContracts: []string{contract},
	}
	resp, err := cli.DescribeReservedContracts(descParams)
	if err != nil {
		logrus.Errorln(err)
	}
	if resp.TotalCount == 0 {
		params.AutoRenew = 0
		params.Months = 0
		return
	}

	respContract := resp.ReservedContractSet[0]
	params.AutoRenew = respContract.AutoRenew
	params.Months = respContract.Months

	return
}
