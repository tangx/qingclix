package global

import (
	"github.com/tangx/qingclix/utils"
	"github.com/tangx/qingyun-sdk-go/qingyun"
)

const (
	qingclixConfigFilename = ".qingclix/config.json"
)

func LoginQingyun() *qingyun.Client {
	authFile := utils.HomeDir() + "/.qingcloud/config.yaml"

	cli := qingyun.NewWithFile(authFile)
	return cli
}

func PresetConfig() string {
	return utils.HomeDir() + "/" + qingclixConfigFilename
}
