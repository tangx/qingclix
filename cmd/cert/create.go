package cert

import (
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/modules"
)

var CertCmdCreate = &cobra.Command{
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
	CertCmdCreate.Flags().StringVarP(&global.CertCrt, "crt", "c", "public.crt", "certificate content file path")
	CertCmdCreate.Flags().StringVarP(&global.CertKey, "key", "k", "private.key", "certificate key file path")
	CertCmdCreate.Flags().StringVarP(&global.CertName, "name", "n", "", "certificate name")
}
