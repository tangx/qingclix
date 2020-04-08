package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/utils"
	"gopkg.in/AlecAivazis/survey.v1"
)

type PreInstance struct {
	Instance map[string]InstanceItem `yaml:"instance,omitempty" json:"instance,omitempty"`
}

type InstanceItem struct {
	ImageId       string `yaml:"image_id,omitempty" json:"image_id,omitempty"`
	CPU           string `yaml:"cpu,omitempty" json:"cpu,omitempty"`
	Memory        string `yaml:"memory,omitempty" json:"memory,omitempty"`
	InstanceClass string `yaml:"instance_class,omitempty" json:"instance_class,omitempty"`
	OsDiskSize    string `yaml:"os_disk_size,omitempty" json:"os_disk_size,omitempty"`
	LoginMode     string `yaml:"login_mode,omitempty" json:"login_mode,omitempty"`
	LoginKeypair  string `yaml:"login_keypair,omitempty" json:"login_keypair,omitempty"`
	InstanceName  string `yaml:"instance_name,omitempty" json:"instance_name,omitempty"`
	Zone          string `yaml:"zone,omitempty" json:"zone,omitempty"`
	Vxnets        string `yaml:"vxnets.1,omitempty" json:"vxnets.1,omitempty"`
}

// ChoosePreinstanllConfig 选择设置配置
func ChoosePreinstanllConfig(pre PreInstance) InstanceItem {
	var option []string
	for k := range pre.Instance {
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
	return pre.Instance[choice]
}

func launchPreinstall(item InstanceItem) {
	data := struct2map(item)
	cli := global.LoginQingyun()

	b, err := cli.Get("RunInstances", data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}

// struct2map 把 struct 转换成 map
func struct2map(item InstanceItem) (data map[string]string) {
	body, _ := json.Marshal(item)
	json.Unmarshal(body, &data)
	return
}

// loadConfig 读取 preinstall.json 配置
func loadPreinstallConfig() PreInstance {
	body, err := ioutil.ReadFile(utils.HomeDir() + "/.qingclix/preinstall.json")
	if err != nil {
		log.Fatal(err)
	}

	var pre = PreInstance{}
	err = json.Unmarshal(body, &pre)
	if err != nil {
		log.Fatal(err)
	}

	return pre
}
