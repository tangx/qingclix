package cert

import (
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/modules"
)

var CertCmdBind = &cobra.Command{
	Use:   "bind",
	Short: "绑定证书到负载均衡监听器",
	Run: func(cmd *cobra.Command, args []string) {
		// todo:
		// 1. bind cert to lbl
		ok := modules.BindCertsToLBListener(global.CertIDs, global.LBListenerID)

		// 2. bind and update
		// 2.1. get lb by lb
		lb := modules.GetLbByLblID(global.LBListenerID)
		// 2.2. update lb
		if (ok && len(lb) != 0) && !global.SkipUpdateLB {
			modules.UpdateLoadBalancers(lb)
		}
	},
}

func init() {
	CertCmdBind.Flags().StringVarP(&global.LBListenerID, "lbl", "", "", "LB Listener to binding")
	CertCmdBind.Flags().StringVarP(&global.CertIDs, "sc", "", "", "Certificate is binding to LB Listener")
	CertCmdBind.Flags().BoolVarP(&global.SkipUpdateLB, "skip-update-lb", "", false, "if true, force to skip update lb")
}
