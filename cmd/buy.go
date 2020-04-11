package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

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
	runInsResp, err := cli.RunInstances(instance)
	if err != nil {
		panic(err)
	}
	fmt.Println(runInsResp.Instances[0])

	// 创建 volume 硬盘
	fmt.Printf("购买硬盘....")
	volume := choiceConfig.Volume
	volume.Zone = instance.Zone // 保证 volume 和 instance 在相同可用区
	volume.VolumeName = instance.InstanceName
	createVolumeResp, err := cli.CreateVolumes(volume)
	if err != nil {
		panic(err)
	}
	fmt.Println(createVolumeResp.Volumes)

	// 绑定 volume 到 instance

	fmt.Println("绑定 volume 到 instance....")
	attachVolParams := types.AttachVolumesRequest{
		Volumes:  createVolumeResp.Volumes,
		Instance: runInsResp.Instances[0],
		Zone:     volume.Zone,
	}
	attachVolumeResp, err := cli.AttachVolumes(attachVolParams)
	if err != nil {
		panic(err)
	}
	fmt.Println(attachVolumeResp)
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
