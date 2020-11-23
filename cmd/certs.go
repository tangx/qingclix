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
	certCmd.AddCommand(certCmdAssociateToLB)
	certCmd.AddCommand(certCmdDelete)
	certCmd.AddCommand(certCmdDisassociteFromLB)

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

var certCmdAssociateToLB = &cobra.Command{
	Use:   "bind",
	Short: "绑定证书到 SLB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bind server certificate")
	},
}

var certCmdDisassociteFromLB = &cobra.Command{
	Use:   "unbind",
	Short: "DisAssocite Certs from Lb listener",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bind server certificate")
	},
}
