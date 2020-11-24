package global

import (
	"fmt"
	"os"

	"github.com/tangx/qingclix/sdk-go/qingyun"
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

// Certs
var (
	CertCrt  string
	CertKey  string
	CertName string
)

var (
	LoadBalanceID string
	LBListenerID  string
	CertIDs       string
)

var (
	SkipUpdateLB bool
)

var (
	LoadBalancerIDs string
)
