package cmd

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/modules"
	"github.com/tangx/qingclix/utils"
	"gopkg.in/AlecAivazis/survey.v1"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().StringVarP(&instance_target, "instance", "i", "", "target instance to clone")

}

var (
	instance_target string
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		cloneInstance()
	},
}

func cloneInstance() {

	if instance_target == "" {
		logrus.Errorf("no target instance clone")
		return
	}

	var item = ClixItem{}
	insResp := modules.DescInstance(instance_target)

	if len(insResp.InstanceSet) != 1 {
		return
	}

	ins := insResp.InstanceSet[0]
	// instance
	// item.Instance.InstanceClass = ins.InstanceClass
	// item.Instance.InstanceName = ins.InstanceName + "_clone"
	// item.Instance.ImageID = ins.Image.ImageID
	// item.Instance.CPU = ins.VcpusCurrent
	// item.Instance.Memory = ins.MemoryCurrent
	// item.Instance.LoginMode = "keypair"
	// item.Instance.LoginKeypair = ins.KeypairIDS[0]
	// item.Instance.Zone = ins.ZoneID
	// item.Instance.OsDiskSize = ins.Extra.OSDiskSize

	item.Instance = Instance{
		InstanceClass: ins.InstanceClass,
		InstanceName:  ins.InstanceName + "_clone",
		ImageID:       ins.Image.ImageID,
		CPU:           ins.VcpusCurrent,
		Memory:        ins.MemoryCurrent,
		LoginMode:     "keypair",
		LoginKeypair:  ins.KeypairIDS[0],
		Zone:          ins.ZoneID,
		OsDiskSize:    ins.Extra.OSDiskSize,
	}

	for _, vxnet := range ins.Vxnets {
		item.Instance.Vxnets = append(item.Instance.Vxnets, vxnet.VxnetID)
	}

	// contract
	contractResp := modules.DescContract(ins.ReservedContract)
	if contractResp.TotalCount == 1 {
		contract := contractResp.ReservedContractSet[0]
		if contract.AutoRenew == 1 {
			item.Contract.AutoRenew = contract.AutoRenew
			item.Contract.Months = 1
		}
	}

	// volume
	for _, volID := range ins.VolumeIDS {
		resp := modules.DescVolume(volID)
		if resp.TotalCount != 1 {
			continue
		}
		vol := resp.DescribeVolumeSet[0]

		volume := Volume{
			Size:       vol.Size,
			VolumeType: vol.VolumeType,
		}
		item.Volumes = append(item.Volumes, volume)
	}

	config := addItem(item.Instance.InstanceName, item)
	dumpConfig(config)
}

func dumpConfig(config ClixConfig) {
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

func addItem(name string, item ClixItem) (config ClixConfig) {
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

func chooseItem() ClixItem {

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
