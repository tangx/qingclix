package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/cmd/cert"
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
	certCmd.AddCommand(cert.CertCmdCreate)
	certCmd.AddCommand(cert.CertCmdDelete)
	certCmd.AddCommand(cert.CertCmdBind)
	certCmd.AddCommand(cert.CertCmdUnbind)
	certCmd.AddCommand(cert.CertCmdSearch)
	certCmd.AddCommand(cert.CertCmdExplore)
}
