package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/utils"
	"gopkg.in/AlecAivazis/survey.v1"
)

// instanceCmd represents the instance command
var instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "服务器实例操作",
	Long: `针对服务器实例进行操作。
例如: 服务器配置管理、 服务器生命管理 等`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(instanceCmd)
}

// LoadPresetConfig 读取预设配置
func LoadPresetConfig() PresetConfig {
	body, err := ioutil.ReadFile(global.ConfigFile)
	logrus.Debugf("%s", body)

	if err != nil {
		ConfigHelp()
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

// ChooseItem 选择预设配置
func ChooseItem(preset PresetConfig) ItemConfig {

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

// 保存配置文件
func SaveConfigToFile(preset PresetConfig) {
	data, err := json.MarshalIndent(preset, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(global.ConfigFile, data, 0644)
	if err != nil {
		logrus.Errorln(err)
	}
}

// InitialConfig 初始化配置
func InitialConfig() {
	base := `{"configs":{}}`

	_, err := utils.MkdirAll(filepath.Dir(global.ConfigFile))

	if err != nil {
		logrus.Infoln(err)
	}

	err = ioutil.WriteFile(global.ConfigFile, []byte(base), 0644)
	if err != nil {
		logrus.Fatalln(err)
	}
}

func ConfigHelp() {
	usage := `Help:
请检查配置文件[%s] 
1. 是否存在
2. Json/YAML 结构是否正确

或使用 qingclix instance configure --initial 初始化
  注意: 该命令会清空、原始数据。


`
	fmt.Printf(usage, global.ConfigFile)
}
