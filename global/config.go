package global

import (
	"fmt"
	"os"

	"github.com/tangx/qingyun-sdk-go/qingyun"
)

var (
	AuthFile = fmt.Sprintf("%s/.qingcloud/config.yaml", os.Getenv("HOME"))

	ConfigFile = fmt.Sprintf("%s/.qingclix/config.json", os.Getenv("HOME"))

	QingtypesFile = fmt.Sprintf("%s/.qingclix/qingtypes.json", os.Getenv("HOME"))
)

var (
	QingClix = qingyun.NewWithFile(AuthFile)
)

// global Flags
var (
	SkipContract bool
	Verbose      int
)
