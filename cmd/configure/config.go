package configure

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/utils"
	"gopkg.in/AlecAivazis/survey.v1"
)

type ClixConfig struct {
	Configs map[string]ClixItem `json:"configs"`
}

type ClixItem struct {
	Instance ItemInstance `json:"instance"`
	Volumes  []ItemVolume `json:"volumes"`
	Contract ItemContract `json:"contract"`
}

type ItemContract struct {
	Months    int `json:"months"`
	AutoRenew int `json:"auto_renew"`
}

type ItemInstance struct {
	ImageID       string   `json:"image_id,omitempty" url:"image_id,omitempty"`
	CPU           int      `json:"cpu,omitempty" url:"cpu,omitempty"`
	Memory        int      `json:"memory,omitempty" url:"memory,omitempty"`
	InstanceClass int      `json:"instance_class,omitempty" url:"instance_class,omitempty"`
	LoginMode     string   `json:"login_mode,omitempty" url:"login_mode,omitempty"`
	LoginKeypair  string   `json:"login_keypair,omitempty" url:"login_keypair,omitempty"`
	InstanceName  string   `json:"instance_name,omitempty" url:"instance_name,omitempty"`
	Zone          string   `json:"zone,omitempty" url:"zone,omitempty"`
	Vxnets        []string `json:"vxnets,omitempty" url:"vxnets,omitempty"`
	OsDiskSize    int      `json:"os_disk_size,omitempty" url:"os_disk_size,omitempty"`
	GPU           int      `json:"gpu,omitempty" url:"gpu,omitempty"`
	GpuClass      int      `json:"gpu_class,omitempty" url:"gpu_class,omitempty"`
	CpuModel      string   `json:"cpu_model,omitempty" url:"cpu_model,omitempty"`
}

type ItemVolume struct {
	Size       int    `json:"size"`
	VolumeType int    `json:"volume_type"`
	VolumeName string `json:"volume_name"`
}

func DumpConfig(config ClixConfig) {
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		logrus.Errorf("%s", err)
	}

	if ok, _ := utils.PathExists(global.ConfigFile); !ok {
		utils.MkdirAll(path.Dir(global.ConfigFile))
	}
	err = ioutil.WriteFile(global.ConfigFile, data, 0644)
	if err != nil {
		logrus.Fatalf("%s", err)
	}
}

func AddItem(name string, item ClixItem) (config ClixConfig) {
	config = LoadConfig()
	config.Configs[name] = item
	return
}

func LoadConfig() (config ClixConfig) {
	ok, err := utils.PathExists(global.ConfigFile)
	if err != nil {
		logrus.Fatalf("%s", err)
	}
	if !ok {
		utils.MkdirAll(path.Dir(global.ConfigFile))
		return
	}

	data, err := ioutil.ReadFile(global.ConfigFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func ChooseItem() ClixItem {

	config := LoadConfig()

	opts := []string{}
	for k := range config.Configs {
		opts = append(opts, k)
	}

	name := ""
	qs := &survey.Select{
		Message: "choose a item name:",
		Options: opts,
	}
	err := survey.AskOne(qs, &name, nil)
	if err != nil {
		logrus.Fatalf("%s", err)
	}

	if name == "" {
		return ClixItem{}
	}

	return config.Configs[name]

}
