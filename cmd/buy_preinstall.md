package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"

	"github.com/tangx/qingclix/resources"
	"github.com/tangx/qingclix/utils"
	"gopkg.in/AlecAivazis/survey.v1"
)

type PreInstance struct {
	// Configs map[string]InstanceItem `yaml:"instance,omitempty" json:"instance,omitempty"`
	Configs map[string]PreConfig `yaml:"instance,omitempty" json:"instance,omitempty"`
}

type PreConfig struct {
	Instance *resources.InstanceRequest `yaml:"instance,omitempty" json:"instance,omitempty"`
	// Volume   *resources.VolumeRequest
}

// LoadPreintallCofnig 读取 config.json 配置
func LoadPreintallCofnig() PreInstance {
	body, err := ioutil.ReadFile(utils.HomeDir() + "/.qingclix/config.json")
	// fmt.Printf("%s", body)
	if err != nil {
		log.Fatal(err)
	}

	var configs = PreInstance{}
	err = json.Unmarshal(body, &configs)
	if err != nil {
		log.Fatal(err)
	}

	return configs
}

// ChoosePreinstanllConfig 选择设置配置
func ChoosePreinstanllConfig(config PreInstance) PreConfig {

	var option []string
	for k := range config.Configs {

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

	return config.Configs[choice]
}
