package cert

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/modules"
)

var CertCmdExplore = &cobra.Command{
	Use:   "explore",
	Short: "查询证书绑定的 LB 和 LBL",
	Run: func(cmd *cobra.Command, args []string) {
		if global.CertIDs == "" {
			_ = cmd.Help()
		}

		explore()
	},
}

func init() {
	CertCmdExplore.Flags().StringVarP(&global.CertIDs, "sc", "", "", "server certificate id")
}

func explore() {
	for _, sc := range strings.Split(global.CertIDs, ",") {
		_ = modules.GetCertBindTo(sc)
	}
}
