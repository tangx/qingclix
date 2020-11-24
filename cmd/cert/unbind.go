package cert

import (
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/modules"
)

var CertCmdUnbind = &cobra.Command{
	Use:   "unbind",
	Short: "DisAssocite Certs from Lb listener",
	Run: func(cmd *cobra.Command, args []string) {
		ok := modules.UnbindCertsFromLBListener(global.CertIDs, global.LBListenerID)
		lb := modules.GetLbByLblID(global.LBListenerID)
		if (ok && len(lb) != 0) && !global.SkipUpdateLB {
			modules.UpdateLoadBalancers(lb)
		}
	},
}

func init() {
	CertCmdUnbind.Flags().StringVarP(&global.LBListenerID, "lbl", "", "", "LB Listener to binding")
	CertCmdUnbind.Flags().StringVarP(&global.CertIDs, "sc", "", "", "Certificate is binding to LB Listener")
	CertCmdUnbind.Flags().BoolVarP(&global.SkipUpdateLB, "skip-update-lb", "", false, "if true, force to skip update lb")
}
