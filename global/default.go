package global

import "github.com/tangx/qingclix/utils"

var (
	HomeDir    string = utils.HomeDir()
	AuthFile   string = HomeDir + "/.qingcloud/config.yaml"
	ConfigFile string = HomeDir + "/.qingclix/config.json"
)

// Global Flags Vars
var (
	// logrus 日志等级
	// 0: Panic (最高), 6: Trace (最低)
	Verbose int

	// 是否跳过 Contract 购买
	SkipContract bool

	// Count 购买数量
	Count int
)
