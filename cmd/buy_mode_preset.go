package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"gopkg.in/AlecAivazis/survey.v1"
)

// presetMode 预设模式
func presetMode() {
	preset := LoadPresetConfig()
	// instance := preset.Configs
	// fmt.Println(preset)
	item := ChooseConfig(preset)

	LaunchInstance(item)

}

// LoadPresetConfig 读取预设配置
func LoadPresetConfig() PresetConfig {
	body, err := ioutil.ReadFile(global.ConfigFile)
	logrus.Debugf("%s", body)

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
