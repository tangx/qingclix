package global

import (
	"github.com/tangx/qingclix/utils"
	"github.com/tangx/qingyun-sdk-go"
)

func LoginQingyun() *qingyun.Client {
	authFile := utils.HomeDir() + "/.qingcloud/config.yaml"

	cli := qingyun.NewWithFile(authFile)
	return cli
}
