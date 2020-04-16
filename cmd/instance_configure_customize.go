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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		QingclixSetLogLevel(8)

		configureCustomizeMain()
	},
}

var (
	configure_customize_qingtypes    string
	configure_customize_dump_defualt bool
)

func init() {
	configureCmd.AddCommand(customizeCmd)
	customizeCmd.Flags().StringVarP(&configure_customize_qingtypes, "qingtype", "", "/path/to/file.json", "使用用户自定义 qingtypes 数据替代默认值。以便适应青云资源调整")

	customizeCmd.Flags().BoolVarP(&configure_customize_dump_defualt, "dump_default", "", false, "打印默认 qingtypes 信息")
}

func configureCustomizeMain() {

	if configure_customize_dump_defualt {
		dumpDefaultQingtypes()
		return
	}

	customizeConfig()

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

func customizeConfig() {
	qingtypes, err := loadQingtypes()
	if err != nil {
		logrus.Error(err)
	}

	insTypes := qingtypes.InstanceType
	imageTypes := qingtypes.ImageType
	logrus.Debug(insTypes)

	insPrarms := customizeInstance(insTypes, imageTypes)
	logrus.Debug(insPrarms)

	volTypes := qingtypes.VolumeType
	volPramas := customizeVolume(volTypes)
	logrus.Debug(volPramas)
}

func customizeInstance(insTypes map[string]types.InstanceType, imageTypes map[string]types.ImageType) (params types.RunInstancesRequest) {

	var inames []string
	for idx := range insTypes {
		logrus.Debug("instance idx=", idx)
		inames = append(inames, idx)
	}

	var images []string
	for idx := range imageTypes {
		logrus.Debug("image idx=", idx)
		images = append(images, idx)
	}

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
	}

	return
}

func customizeVolume(volTypes map[string]types.VolumeType) (params types.CreateVolumesRequest) {
	var names []string
	for idx := range volTypes {
		logrus.Debug("instance idx=", idx)
		names = append(names, idx)
	}

	var qs = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Select{
				Message: "选择磁盘类型",
				Options: names,
			},
		},
		{
			Name: "size",
			Prompt: &survey.Select{
				Message: "选择磁盘大小",
				Options: []string{"100", "200", "300", "500", "1000", "2000"},
			},
		},
	}

	answers := struct {
		Name string
		Size int
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Debug(answers)

	params = types.CreateVolumesRequest{
		Size:       answers.Size,
		VolumeType: volTypes[answers.Name].Type,
	}
	return
}
