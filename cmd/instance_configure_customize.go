package cmd

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/types"
	"gopkg.in/AlecAivazis/survey.v1"
)

// customizeCmd represents the customize command
var customizeCmd = &cobra.Command{
	Use:   "customize",
	Short: "使用交互界面，自定义配置",
	Long:  `使用交互界面，灵活生成所选配置`,
	Run: func(cmd *cobra.Command, args []string) {
		QingclixSetLogLevel(global.Verbose)

		configureCustomizeMain()
	},
}

var (
	configure_customize_qingtypes    string
	configure_customize_dump_defualt bool
	configure_customize_label        string
)

func init() {
	configureCmd.AddCommand(customizeCmd)
	customizeCmd.Flags().StringVarP(&configure_customize_qingtypes, "qingtype", "", "/path/to/file.json", "使用用户自定义 qingtypes 数据替代默认值。以便适应青云资源调整")
	customizeCmd.Flags().BoolVarP(&configure_customize_dump_defualt, "dump_default", "", false, "打印默认 qingtypes 信息")
	customizeCmd.Flags().StringVarP(&configure_customize_label, "label", "l", "customize-item", "增加配置的名称")

}

func configureCustomizeMain() {

	if configure_customize_dump_defualt {
		dumpDefaultQingtypes()
		return
	}
	item := customizeConfig()
	appendItemConfig(configure_customize_label, item)
}

func dumpDefaultQingtypes() {
	fmt.Println(global.QingTypes)
}

func loadQingtypes() (qtypes types.Qingtypes, err error) {
	var data []byte

	if configure_customize_qingtypes != `/path/to/file.json` {
		data, err = ioutil.ReadFile(configure_customize_qingtypes)
		if err != nil {
			logrus.Debug(err)
			return
		}
	} else {
		data = []byte(global.QingTypes)
	}

	// 覆盖正常流程值
	var manual_debuger bool = true
	if manual_debuger {
		data, _ = ioutil.ReadFile("/Users/tangxin/.qingclix/qingtypes.json")
	}

	logrus.Debug(string(data))
	return types.LoadQingTypes(data)
}

// 根据 Instance Class ID 查询 Instance Type ID 值
func instanceClassToType(class int, instypes []types.InstanceType) (instype string) {

	for _, instype := range instypes {
		if instype.Class == class {
			return instype.Type
		}
	}
	return
}

func customizeConfig() ItemConfig {
	qingtypes, err := loadQingtypes()
	if err != nil {
		logrus.Error(err)
	}

	insTypes := qingtypes.InstanceTypes
	logrus.Debug(insTypes)

	insPrarms := customizeInstance(qingtypes)
	insPrarms.InstanceName = configure_customize_label
	logrus.Debug(insPrarms)

	// get instance type
	instype := instanceClassToType(insPrarms.InstanceClass, qingtypes.InstanceTypes)

	// // get all volumesType
	// volTypes := qingtypes.VolumeTypes
	// // get support volumesType
	// volTypes := supportVolumeType(instype, qingtypes.VolumeTypes)

	support := qingtypes.Relation[instype]
	volTypes := supportVolumeType2(instype, support, qingtypes.VolumeTypes)

	volsPramas := customizeVolume(volTypes)
	logrus.Debug(volsPramas)

	contractParams := customizeContract()
	logrus.Debug(contractParams)

	return ItemConfig{
		Instance: insPrarms,
		Volumes:  volsPramas,
		Contract: contractParams,
	}
}

func customizeInstance(qingtypes types.Qingtypes) (params types.RunInstancesRequest) {

	var qsOpts = make(map[string]types.InstanceType)
	var instances []string
	for _, instype := range qingtypes.InstanceTypes {
		tip := fmt.Sprintf("%s [%s] -- %s", instype.Name, instype.Type, instype.Desc)
		qsOpts[tip] = instype
		instances = append(instances, tip)
	}
	logrus.Debug(instances)

	// image
	var qsImageOpts = make(map[string]types.ImageType)
	var images []string
	for _, imgtype := range qingtypes.ImageTypes {
		tip := fmt.Sprintf("%s [%s] -- %s", imgtype.Name, imgtype.Image, imgtype.Desc)
		qsImageOpts[tip] = imgtype
		images = append(images, tip)
	}
	logrus.Debug("images = ", images)
	logrus.Debug("qsImageOpts = ", qsImageOpts)

	// vxnets
	var qsVxnetOpts = make(map[string]types.CommonType)
	var vxnets []string
	for _, commontype := range qingtypes.Vxnets {
		tip := fmt.Sprintf("%s -- %s", commontype.Name, commontype.Desc)
		qsVxnetOpts[tip] = commontype
		vxnets = append(vxnets, tip)
	}
	logrus.Debug(qsVxnetOpts)

	// zones
	var qsZoneOpts = make(map[string]types.CommonType)
	var zones []string
	for _, commontype := range qingtypes.Zones {
		tip := fmt.Sprintf("%s -- %s", commontype.Name, commontype.Desc)
		qsZoneOpts[tip] = commontype
		zones = append(zones, tip)
	}
	logrus.Debug("zones = ", zones)
	logrus.Debug("qsZoneOpts = ", qsZoneOpts)

	// keypairs
	var qsKeypairOpts = make(map[string]types.CommonType)
	var keypairs []string
	for _, commontype := range qingtypes.Keypairs {
		tip := fmt.Sprintf("%s -- %s", commontype.Name, commontype.Desc)
		qsKeypairOpts[tip] = commontype
		keypairs = append(keypairs, tip)
	}
	logrus.Debug(keypairs)

	// ask question
	var qs = []*survey.Question{
		{
			Name: "instype",
			Prompt: &survey.Select{
				Message: "选择服务器类型",
				Options: instances,
			},
		},
		{
			Name: "cpu",
			Prompt: &survey.Select{
				Message: "选择 CPU 核数",
				Options: []string{"1", "2", "4", "8", "16", "32", "64"},
				Default: "2",
			},
		},
		{
			Name: "ratio",
			Prompt: &survey.Select{
				Message: "选择 cpu:memory 比例",
				Options: []string{"1:2", "1:4"},
				Default: "1:2",
			},
		},
		{
			Name: "image",
			Prompt: &survey.Select{
				Message: "选择操作系统镜像",
				Options: images,
			},
		},
		{
			Name: "zone",
			Prompt: &survey.Select{
				Message: "选择可用区",
				Options: zones,
			},
		},
		{
			Name: "vxnet",
			Prompt: &survey.Select{
				Message: "选择网络",
				Options: vxnets,
			},
		},
		{
			Name: "keypair",
			Prompt: &survey.Select{
				Message: "选择登录密钥",
				Options: keypairs,
			},
		},
	}

	answers := struct {
		Instype string
		Ratio   string
		CPU     int
		Memory  int
		Image   string
		Zone    string
		Vxnet   string
		Keypair string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Debug(answers)

	// 倍率转换
	ratio, _ := strconv.Atoi(strings.Split(answers.Ratio, ":")[1])

	logrus.Debug("answers.Name = ", answers.Zone)
	logrus.Debug("qsZoneOpts[answers.Zone] = ", qsZoneOpts[answers.Zone])

	// transfer
	params = types.RunInstancesRequest{
		InstanceClass: qsOpts[answers.Instype].Class,
		CPU:           answers.CPU,
		Memory:        answers.CPU * 1024 * ratio,
		ImageID:       qsImageOpts[answers.Image].Image,
		LoginKeypair:  qsKeypairOpts[answers.Keypair].Name,
		LoginMode:     "keypair",
		Zone:          qsZoneOpts[answers.Zone].Name,
		Vxnets:        []string{qsVxnetOpts[answers.Vxnet].Name},
	}

	return
}

func supportVolumeType2(instancetype string, support []int, volTypes []types.VolumeType) (support_types []types.VolumeType) {

	// var support []int
	// switch instancetype {
	// case "e1", "e2", "p1":
	// 	support = []int{2, 3, 5, 200}

	// case "s1":
	// 	support = []int{0, 100}
	// }

	for _, v := range support {
		for _, v2 := range volTypes {
			if v == v2.Type {
				support_types = append(support_types, v2)
			}
		}
	}
	return
}

func customizeVolume(volTypes []types.VolumeType) (params []types.CreateVolumesRequest) {

	var qsOpts = make(map[string]types.VolumeType)
	var volnames []string
	for _, voltype := range volTypes {
		tip := fmt.Sprintf("%s -- %s", voltype.Name, voltype.Desc)

		qsOpts[tip] = voltype
		volnames = append(volnames, tip)
	}
	logrus.Debug(volnames)
	logrus.Debug(qsOpts)

	answers := struct {
		Name string
		Size int
	}{}

	var qs = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Select{
				Message: "选择磁盘类型",
				Options: volnames,
			},
		},
		{
			Name: "size",
			Prompt: &survey.Select{
				Message: "选择磁盘大小",
				Options: []string{"100", "200", "300", "500", "1000", "2000"},
				Default: "100",
			},
		},
	}

	for i := 1; i <= 4; i++ {
		var qsAddVolume = []*survey.Question{
			{
				Name: "continue",
				Prompt: &survey.Select{
					Message: fmt.Sprintf("添加硬盘 %d/4 ", i),
					Options: []string{"true", "false"},
					Default: "true",
				},
			},
		}

		var next bool
		survey.Ask(qsAddVolume, &next)
		if !next {
			break
		}

		survey.Ask(qs, &answers)
		logrus.Debug(answers)
		logrus.Debug("answers.Name = ", answers.Name)
		logrus.Debug("qsOpts[answers.Name] = ", qsOpts[answers.Name])

		valumeParams := types.CreateVolumesRequest{
			Size:       answers.Size,
			VolumeType: qsOpts[answers.Name].Type,
		}
		params = append(params, valumeParams)

	}
	return
}

func customizeContract() (params types.ApplyReservedContractWithResourcesRequest) {
	var qsMonths = []*survey.Question{
		{
			Name: "months",
			Prompt: &survey.Select{
				Message: "是否包月",
				Options: []string{"true", "false"},
				Default: "true",
			},
		},
	}
	var qsAutoRenew = []*survey.Question{
		{
			Name: "autoRenew",
			Prompt: &survey.Select{
				Message: "是否自动续费（true）",
				Options: []string{"true", "false"},
				Default: "true",
			},
		},
	}

	answers := struct {
		AutoRenew bool
		Months    bool
	}{}

	err := survey.Ask(qsMonths, &answers)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Debug(answers)

	if answers.Months {
		params.Months = 1
		survey.Ask(qsAutoRenew, &answers)
	}
	if answers.AutoRenew {
		params.AutoRenew = 1
	}

	return
}
