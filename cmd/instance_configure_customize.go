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

func customizeConfig() ItemConfig {
	qingtypes, err := loadQingtypes()
	if err != nil {
		logrus.Error(err)
	}

	insTypes := qingtypes.InstanceTypes
	imageTypes := qingtypes.ImageTypes
	logrus.Debug(insTypes)

	insPrarms := customizeInstance(insTypes, imageTypes)
	insPrarms.InstanceName = configure_customize_label
	logrus.Debug(insPrarms)

	volTypes := qingtypes.VolumeTypes
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

func customizeInstance(insTypes map[string]types.InstanceType, imageTypes map[string]types.ImageType) (params types.RunInstancesRequest) {

	var inames []string
	for idx := range insTypes {
		inames = append(inames, idx)
	}
	logrus.Debug(inames)

	var images []string
	for idx := range imageTypes {
		images = append(images, idx)
	}
	logrus.Debug(images)

	// ask question
	var qs = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Select{
				Message: "选择服务器类型",
				Options: inames,
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
				Default: "2",
			},
		},
	}

	answers := struct {
		Name   string
		Ratio  string
		CPU    int
		Memory int
		Image  string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Debug(answers)

	// 倍率转换
	ratio, _ := strconv.Atoi(strings.Split(answers.Ratio, ":")[1])

	// transfer
	params = types.RunInstancesRequest{
		InstanceClass: insTypes[answers.Name].Class,
		CPU:           answers.CPU,
		Memory:        answers.CPU * 1024 * ratio,
		ImageID:       answers.Image,
		LoginKeypair:  "kp-2kodyll8",
		LoginMode:     "keypair",
	}

	return
}

func customizeVolume(volTypes map[string]types.VolumeType) (params []types.CreateVolumesRequest) {
	var names []string
	for idx := range volTypes {
		names = append(names, idx)
	}
	logrus.Debug(names)

	answers := struct {
		Name     string
		Size     int
		Continue bool
	}{}

	var qs = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Select{
				Message: "选择磁盘类型",
				Options: names,
				Default: "200",
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
		survey.Ask(qsAddVolume, &answers)
		if !answers.Continue {
			break
		}

		survey.Ask(qs, &answers)
		logrus.Debug(answers)

		valumeParams := types.CreateVolumesRequest{
			Size:       answers.Size,
			VolumeType: volTypes[answers.Name].Type,
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
