package cmd

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/types"
)

// installCmd represents the buy command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "根据预设信息购买机器",
	Long: `根据预设信息购买机器
  1. 根据 ~/.qingclix/config.json 预设信息，直接选择购买对应的机器
  2. 使用 -i 进入交互界面，选择自定购买参数`,
	Run: func(cmd *cobra.Command, args []string) {
		// logrus.SetLevel(logrus.TraceLevel)
		SetLogLevel(global.Verbose)

		instanceInstallMain()
	},
}

const (
	ModeDescription = `服务器购买模式
clone: 克隆一台已有机器
interactive: 交互提示购买
preset: 预设值模式, 从 ~/.qingclix/config.json 中读取预设文件`
)

func init() {
	instanceCmd.AddCommand(installCmd)
}

type PresetConfig struct {
	Configs map[string]ItemConfig `yaml:"configs,omitempty" json:"configs,omitempty"`
}
type ItemConfig struct {
	Instance types.RunInstancesRequest                       `yaml:"instance,omitempty" json:"instance,omitempty"`
	Volumes  []types.CreateVolumesRequest                    `yaml:"volumes,omitempty" json:"volumes,omitempty"`
	Contract types.ApplyReservedContractWithResourcesRequest `yaml:"contract,omitempty" json:"contract,omitempty"`
}

// launch 创建实例
func instanceInstallMain() {

	preset := LoadPresetConfig()
	item := ChooseItem(preset)

	InstallInstance(item)
}

// InstallInstance 根据配置，购买服务器、磁盘及合约, 外层调用
func InstallInstance(config ItemConfig) {

	// Inital a Client
	cli := types.Client{}

	// 直接创建
	if global.Count == 0 {

		installInstance(cli, config)
		return
	}

	// 循环模式
	originName := config.Instance.InstanceName
	for i := 0; i < global.Count; i++ {
		PostfixNumber := i + PostfixStartNumber
		config.Instance.InstanceName = fmt.Sprintf("%s-%d", originName, PostfixNumber)

		logrus.Debugln("config.Instance.InstanceName = ", config.Instance.InstanceName)

		installInstance(cli, config)
	}
}

// installInstance 根据配置， 购买服务器、 磁盘及合约
func installInstance(cli types.Client, config ItemConfig) {

	if global.Dryrun {
		logrus.Infoln("dryrun 模式, 跳过购买, 直接返回")
		return
	}

	// 购买主机
	instanceConfig := config.Instance
	instances := buyInstance(cli, instanceConfig)

	// 购买磁盘

	// 定义磁盘集合 如果一台机器有多块磁盘
	var volumesSet [][]string

	for _, config := range config.Volumes {
		logrus.Debugln(config)

		config.Zone = instanceConfig.Zone // 保证 volume 和 instance 在相同可用区
		if config.VolumeName == "" {
			config.VolumeName = instanceConfig.InstanceName
		}
		volumes := buyVolumeForInstance(cli, instances[0], config)
		// 绑定磁盘到主机
		attachResult := attachVolumeToInstance(cli, instances[0], volumes, instanceConfig.Zone)
		if attachResult {
			logrus.Infof("attach Volimes(%s) to Instance(%s): %t", volumes, instances[0], attachResult)
		}

		volumesSet = append(volumesSet, volumes)
	}

	// 资源付费
	contractConfig := config.Contract
	contractConfig.Zone = instanceConfig.Zone
	// 付费服务器
	payResources(cli, instances, contractConfig)
	// 付费硬盘
	for _, volumes := range volumesSet {
		payResources(cli, volumes, contractConfig)
	}

	fmt.Println("")
}

// applyLeaseAssociateContract 根据目标资源申请购买、支付合约，并绑定到对应的资源上。
func payResources(cli types.Client, resources []string, params types.ApplyReservedContractWithResourcesRequest) (ok bool) {

	if global.SkipContract {
		logrus.Infof("强制跳过 %s 合约过程", resources)
		return false
	}

	params.Resources = resources

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
		logrus.Errorln(leaseResp.Message)
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
func buyInstance(cli types.Client, instance types.RunInstancesRequest) (instances []string) {

	// 创建 instance 实例
	fmt.Printf("购买主机....")
	runResp, err := cli.RunInstances(instance)
	if err != nil {
		panic(err)
	}
	fmt.Println(runResp)

	return runResp.Instances
}

// buyVolumeForInstance 为主机购买硬盘
// 注意 可用区 与 磁盘类型 匹配
func buyVolumeForInstance(cli types.Client, instance string, volume types.CreateVolumesRequest) (volumes []string) {

	// 创建 volume 硬盘
	fmt.Printf("购买硬盘....")
	createResp, err := cli.CreateVolumes(volume)
	if err != nil {
		panic(err)
	}
	fmt.Println(createResp)

	return createResp.Volumes
}

// payResources 对传入资源创建预留合约，并付费
// func payResources(resources []string, months int, auto_renew int) (ok bool) {
// 	return
// }

// attachVolumeToInstance 挂载磁盘到主机
func attachVolumeToInstance(cli types.Client, instance string, volumes []string, zone string) (ok bool) {

	// 判断服务器是否为可用状态
	fmt.Printf("等待服务器是否为 running 状态 ..")
	descInstanceParams := types.DescribeInstancesRequest{
		Instances: []string{instance},
		Status:    []string{"running"},
		Zone:      zone,
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

	// 判断 volume 的状态是否为需求状态。
	// ex:
	// func IsReadyVolume(volume string,status string) (ok bool) {}
	fmt.Printf("等待磁盘是否为 available 状态 ..")
	descParams := types.DescribeVolumesRequest{
		Volumes: volumes,
		Status:  []string{"available"},
		Zone:    zone,
	}
	for {
		resp, _ := cli.DescribeVolumes(descParams)
		if resp.TotalCount == len(descParams.Volumes) {
			fmt.Println(".. OK")
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Printf(".")
	}

	// 绑定 volume 到 instance
	fmt.Printf("绑定 volume 到 instance....")
	attachVolParams := types.AttachVolumesRequest{
		Volumes:  volumes,
		Instance: instance,
		Zone:     zone,
	}
	attachVolumeResp, err := cli.AttachVolumes(attachVolParams)
	if err != nil {
		logrus.Errorln(err)
		return false
	}
	fmt.Println(attachVolumeResp)
	return true
}
