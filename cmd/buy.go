package cmd

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/types"
)

// buyCmd represents the buy command
var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "根据预设信息购买机器",
	Long: `根据预设信息购买机器
  1. 根据 ~/.qingclix/preinstall.json 预设信息，直接选择购买对应的机器
  2. 使用 -i 进入交互界面，选择自定购买参数`,
	Run: func(cmd *cobra.Command, args []string) {
		launch()
	},
}

// Flags 变量
var (
	// 执行模式
	mode string
)

const (
	ModeDesc = `服务器购买模式
clone: 克隆一台已有机器
interactive: 交互提示购买
preset: 预设值模式, 从 ~/.qingclix/config.json 中读取预设文件`
)

func init() {
	rootCmd.AddCommand(buyCmd)
	buyCmd.Flags().StringVarP(&mode, "mode", "m", "preset", ModeDesc)
}

type PresetConfig struct {
	Configs map[string]ItemConfig `yaml:"configs,omitempty" json:"configs,omitempty"`
}
type ItemConfig struct {
	Instance types.RunInstancesRequest                       `yaml:"instance,omitempty" json:"instance,omitempty"`
	Volume   types.CreateVolumesRequest                      `yaml:"volume,omitempty" json:"volume,omitempty"`
	Contract types.ApplyReservedContractWithResourcesRequest `yaml:"contract,omitempty" json:"contract,omitempty"`
}

// launch 创建实例
func launch() {
	switch mode {
	case "preset":
		presetMode()
	// case "clone":
	// 	cloneMode()
	// case "interactive":
	// 	interactiveMode()
	default:
		presetMode()
	}

}

// launchInstance 根据配置，购买服务器、磁盘及合约
func launchInstance(itemConfig ItemConfig) {

	// Inital a Client
	cli := types.Client{}

	// 创建 instance 实例
	fmt.Printf("购买主机....")
	instance := itemConfig.Instance
	runInstanceResp, err := cli.RunInstances(instance)
	if err != nil {
		panic(err)
	}
	fmt.Println(runInstanceResp)

	// 判断服务器是否为可用状态
	fmt.Printf("等待服务器是否为 running 状态 ..")
	descInstanceParams := types.DescribeInstancesRequest{
		Zone:      instance.Zone,
		Instances: runInstanceResp.Instances,
		Status:    []string{"running"},
	}
	for {
		resp, _ := cli.DescribeInstances(descInstanceParams)
		if resp.TotalCount == len(descInstanceParams.Instances) {
			fmt.Println(".. OK")
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Printf(".")
	}

	// 创建 volume 硬盘
	fmt.Printf("购买硬盘....")
	volume := itemConfig.Volume
	volume.Zone = instance.Zone // 保证 volume 和 instance 在相同可用区
	volume.VolumeName = instance.InstanceName
	createVolumeResp, err := cli.CreateVolumes(volume)
	if err != nil {
		panic(err)
	}
	fmt.Println(createVolumeResp)

	// 判断 volume 的状态是否为需求状态。
	// ex:
	// func IsReadyVolume(volume string,status string) (ok bool) {}
	fmt.Printf("等待磁盘是否为 available 状态 ..")
	descVolumeParams := types.DescribeVolumesRequest{
		Zone:    instance.Zone,
		Volumes: createVolumeResp.Volumes,
		Status:  []string{"available"},
	}
	for {
		resp, _ := cli.DescribeVolumes(descVolumeParams)
		if resp.TotalCount == len(descVolumeParams.Volumes) {
			fmt.Println(".. OK")
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Printf(".")
	}

	// 绑定 volume 到 instance
	fmt.Printf("绑定 volume 到 instance....")
	attachVolParams := types.AttachVolumesRequest{
		Volumes:  createVolumeResp.Volumes,
		Instance: runInstanceResp.Instances[0],
		Zone:     volume.Zone,
	}
	attachVolumeResp, err := cli.AttachVolumes(attachVolParams)
	if err != nil {
		panic(err)
	}
	fmt.Println(attachVolumeResp)

	// todo: 根据服务器配置
	// 1. 创建 Contract
	// 2. 付费 Contract
	// 3. 绑定 服务器到 Contract

	// instanceContract := contract
	itemConfig.Contract.Zone = instance.Zone

	// 付费服务器
	instanceContract := itemConfig.Contract
	instanceContract.Resources = runInstanceResp.Instances
	ApplyLeaseAssociateContract(cli, instanceContract)
	// 付费硬盘
	volumeContract := itemConfig.Contract
	volumeContract.Resources = createVolumeResp.Volumes
	ApplyLeaseAssociateContract(cli, volumeContract)

}

// ApplyLeaseAssociateContract 根据目标资源申请购买、支付合约，并绑定到对应的资源上。
func ApplyLeaseAssociateContract(cli types.Client, params types.ApplyReservedContractWithResourcesRequest) (ok bool) {

	// 购买合约
	fmt.Printf("购买 %s 的合约: .. ", params.Resources)
	resp, err := cli.ApplyReservedContractWithResources(params)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(resp.ContractID)

	// 支付合约
	fmt.Printf("支付合约: %s ..", resp.ContractID)
	leaseParams := types.LeaseReservedContractRequest{
		Contract: resp.ContractID,
	}
	leaseResp, err := cli.LeaseReservedContract(leaseParams)
	if leaseResp.RetCode == 0 {
		fmt.Println("..OK")
	} else {
		fmt.Errorf("%s\n", leaseResp.Message)
	}

	// 判断合约是否可用
	fmt.Printf("判断合约状态 ...")
	descParams := types.DescribeReservedContractsRequest{
		ReservedContracts: []string{resp.ContractID},
		Zone:              params.Zone,
		Status:            []string{"active"},
	}
	for {
		result, err := cli.DescribeReservedContracts(descParams)
		if err != nil {
			logrus.Infoln("err")
		}

		if result.TotalCount == len(descParams.ReservedContracts) {
			fmt.Println("..OK ")
			break
		}
		fmt.Printf(".")
		time.Sleep(1 * time.Second)
	}

	// 判定合约到目标资源
	fmt.Printf("绑定资源[%s] 到 合约[%s] ...", params.Resources, resp.ContractID)
	assoParams := types.AssociateReservedContractRequest{
		Contract:  resp.ContractID,
		Resources: params.Resources,
	}
	_, err = cli.AssociateReservedContract(assoParams)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("..OK")
	return true

}

// buyInstance 购买主机实例
func buyInstance(instance types.RunInstancesRequest) (instances []string) {
	return
}

// buyVolumeForInstance 为主机购买硬盘
// 注意 可用区 与 磁盘类型 匹配
func buyVolumeForInstance(instance string, volume_size int) (volumes []string) {
	return
}

// payResources 对传入资源创建预留合约，并付费
func payResources(resources []string, months int, auto_renew int) (ok bool) {
	return
}
