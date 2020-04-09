package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/resources"
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

// launch 创建实例
func launch() {
	presetMode()
}

// presetMode 预设模式
func presetMode() {
	preset := LoadPresetConfig()
	config := ChooseConfig(preset)
	// fmt.Println(config)
	RunInstance(config.Instance)
}

type Preset struct {
	Configs map[string]Config `yaml:"configs,omitempty" json:"configs,omitempty"`
}

type Config struct {
	Instance resources.InstanceRequest `yaml:"instance,omitempty" json:"instance,omitempty"`
	Volume   resources.VolumeRequest   `yaml:"volume,omitempty" json:"volume,omitempty"`
}

// LoadPresetConfig 读取预设配置
func LoadPresetConfig() Preset {
	body, err := ioutil.ReadFile(global.PresetConfig())
	if err != nil {
		logrus.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	var preset Preset
	err = json.Unmarshal(body, &preset)
	if err != nil {
		logrus.Fatal(err)
	}
	// fmt.Println(preset)

	return preset
}

// ChooseConfig 选择预设配置
func ChooseConfig(preset Preset) Config {

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

func RunInstance(config resources.InstanceRequest) {
	action := "RunInstances"
	values, err := query.Values(config)
	if err != nil {
		logrus.Fatal("query.Values=", err)
	}
	// fmt.Println(values)

	// fmt.Println("RunInstance", action)
	cli := global.LoginQingyun()
	body, err := cli.GetByUrlValues(action, values)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("%s\n", body)

}
