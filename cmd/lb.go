package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/modules"
)

// certCmd represents the cert command
var lbCmd = &cobra.Command{
	Use:   "lb",
	Short: "负载均衡",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("cert called")
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(lbCmd)
}

var lbCmdUpdate = &cobra.Command{
	Use:   "update",
	Short: "更新负载均衡配置",
	Run: func(cmd *cobra.Command, args []string) {
		if global.Loadbalancers == "" {
			_ = cmd.Help()
			os.Exit(1)
		}
		modules.UpdateLoadBalancers(global.Loadbalancers)
	},
}

func init() {
	lbCmd.AddCommand(lbCmdUpdate)
	lbCmdUpdate.Flags().StringVarP(&global.Loadbalancers, "lb", "", "", "Loadbalances , split with comma (ex lb-123,lb-223)")
}
