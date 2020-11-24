package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/modules"
)

// certCmd represents the cert command
var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "证书管理",

	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("cert called")
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(certCmd)
	certCmd.AddCommand(certCmdCreate)
	certCmd.AddCommand(certCmdAssociateToLBListener)
	certCmd.AddCommand(certCmdDelete)
	certCmd.AddCommand(certCmdDisassociteFromLBListener)
}

var certCmdCreate = &cobra.Command{
	Use:   "create",
	Short: "创建证书",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("create server certificate")
		if global.CertName == "" {
			_ = cmd.Help()
		}

		modules.CreateCertficate(global.CertName, global.CertKey, global.CertCrt)
	},
}

func init() {
	certCmdCreate.Flags().StringVarP(&global.CertCrt, "crt", "c", "public.crt", "certificate content file path")
	certCmdCreate.Flags().StringVarP(&global.CertKey, "key", "k", "private.key", "certificate key file path")
	certCmdCreate.Flags().StringVarP(&global.CertName, "name", "n", "", "certificate name")
}

var certCmdDelete = &cobra.Command{
	Use:   "delete",
	Short: "删除证书",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create server certificate")
	},
}

var certCmdAssociateToLBListener = &cobra.Command{
	Use:   "bind",
	Short: "绑定证书到负载均衡监听器",
	Run: func(cmd *cobra.Command, args []string) {
		// todo:
		// 1. bind cert to lbl
		ok := modules.BindCertsToLBListener(global.BindingCerts, global.BindingLBListener)

		// 2. bind and update
		// 2.1. get lb by lb
		lb := modules.GetLbByLbl(global.BindingLBListener)
		// 2.2. update lb
		if (ok && len(lb) != 0) && !global.SkipUpdateLB {
			modules.UpdateLoadBalancers(lb)
		}
	},
}

func init() {
	certCmdAssociateToLBListener.Flags().StringVarP(&global.BindingLBListener, "lbl", "", "", "LB Listener to binding")
	certCmdAssociateToLBListener.Flags().StringVarP(&global.BindingCerts, "sc", "", "", "Certificate is binding to LB Listener")
	certCmdAssociateToLBListener.Flags().BoolVarP(&global.SkipUpdateLB, "skip-update-lb", "", false, "if true, force to skip update lb")
}

var certCmdDisassociteFromLBListener = &cobra.Command{
	Use:   "unbind",
	Short: "DisAssocite Certs from Lb listener",
	Run: func(cmd *cobra.Command, args []string) {
		ok := modules.UnbindCertsFromLBListener(global.BindingCerts, global.BindingLBListener)
		lb := modules.GetLbByLbl(global.BindingLBListener)
		if (ok && len(lb) != 0) && !global.SkipUpdateLB {
			modules.UpdateLoadBalancers(lb)
		}
	},
}

func init() {
	certCmdDisassociteFromLBListener.Flags().StringVarP(&global.BindingLBListener, "lbl", "", "", "LB Listener to binding")
	certCmdDisassociteFromLBListener.Flags().StringVarP(&global.BindingCerts, "sc", "", "", "Certificate is binding to LB Listener")
	certCmdDisassociteFromLBListener.Flags().BoolVarP(&global.SkipUpdateLB, "skip-update-lb", "", false, "if true, force to skip update lb")
}
