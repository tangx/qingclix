package global

// Config Const
const (
	ConfigFile = "/Users/tangxin/.qingclix/config.json"
	AuthFile   = "/Users/tangxin/.qingcloud/config.yaml"
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
