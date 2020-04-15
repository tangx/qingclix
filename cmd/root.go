package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/utils"
	"gopkg.in/AlecAivazis/survey.v1"
)

var rootCmd = &cobra.Command{
	Use:   "qingclix",
	Short: "qingclix 命令行用来聚合青云控制台的复杂处理流程，简化日常操作",
	Long: `青云控制台的操作常用操作复杂，
  例如，新购机器、更换操作系统 等
  实现目标根据预设参数或配置，快速实现日常操作`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&global.Verbose, "verbose", "v", 4, "logrus 日志等级。 0: Panic, 4: Info, 6: Trace. ")
	rootCmd.PersistentFlags().BoolVarP(&global.SkipContract, "skip_contract", "", false, "强制跳过合约购买过程。")
	rootCmd.PersistentFlags().IntVarP(&global.Count, "count", "c", 1, "设置购买数量")
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func SetLogLevel(level int) {
	logLevel := logrus.Level(level)
	logrus.SetLevel(logLevel)
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
