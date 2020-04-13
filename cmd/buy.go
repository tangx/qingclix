package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/types"
	"gopkg.in/AlecAivazis/survey.v1"
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

func init() {
	rootCmd.AddCommand(buyCmd)
	// buyCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "交互模式")
}

type PresetConfig struct {
	Configs map[string]ItemConfig `yaml:"configs,omitempty" json:"configs,omitempty"`
}
type ItemConfig struct {
	Instance types.RunInstancesRequest  `yaml:"instance,omitempty" json:"instance,omitempty"`
	Volume   types.CreateVolumesRequest `yaml:"volume,omitempty" json:"volume,omitempty"`
	Contract types.CreateVolumesRequest `yaml:"contract,omitempty" json:"contract,omitempty"`
}

// launch 创建实例
func launch() {
	presetMode()
}

// presetMode 预设模式
func presetMode() {

	cli := types.Client{}

	preset := LoadPresetConfig()
	// instance := preset.Configs
	// fmt.Println(preset)
	choiceConfig := ChooseConfig(preset)
	// contract := choiceConfig.Contract

	// 创建 instance 实例
	fmt.Printf("购买主机....")
	instance := choiceConfig.Instance
	runInstanceResp, err := cli.RunInstances(instance)
	if err != nil {
		panic(err)
	}
	fmt.Println(runInstanceResp)
	// fmt.Println(runInstanceResp.Instances)

	// 创建 volume 硬盘
	fmt.Printf("购买硬盘....")
	volume := choiceConfig.Volume
	volume.Zone = instance.Zone // 保证 volume 和 instance 在相同可用区
	volume.VolumeName = instance.InstanceName
	createVolumeResp, err := cli.CreateVolumes(volume)
	if err != nil {
		panic(err)
	}
	fmt.Println(createVolumeResp)

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
}

// LoadPresetConfig 读取预设配置
func LoadPresetConfig() PresetConfig {
	body, err := ioutil.ReadFile(global.ConfigFile)
	if err != nil {
		logrus.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	var preset PresetConfig
	err = json.Unmarshal(body, &preset)
	if err != nil {
		logrus.Fatal(err)
	}
	// fmt.Println(preset)

	return preset
}

// ChooseConfig 选择预设配置
func ChooseConfig(preset PresetConfig) ItemConfig {

	var option []string
	for k := range preset.Configs {
		option = append(option, k)
	}
	// 结果排序，优化展示效果
	sort.Strings(option)

	// 选择
	var qs = []*survey.Question{
		{
			Name: "choice",
			Prompt: &survey.Select{
				Message: "选择购买配置: ",
				Options: option,
			},
		},
	}
	var choice string
	err := survey.Ask(qs, &choice)
	if err != nil {
		log.Fatal(err)
	}

	return preset.Configs[choice]
}
