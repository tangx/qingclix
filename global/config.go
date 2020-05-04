package global

import (
	"fmt"
	"os"

	"github.com/tangx/qingyun-sdk-go/qingyun"
)

var (
	authFile = fmt.Sprintf("%s/.qingcloud/config.yaml", os.Getenv("HOME"))
	Clix     = qingyun.NewWithFile(authFile)
)
