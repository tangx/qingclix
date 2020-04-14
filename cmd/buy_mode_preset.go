package cmd

import (
	"log"
	"sort"

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
