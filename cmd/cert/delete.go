package cert

import (
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/modules"
)

var CertCmdDelete = &cobra.Command{
	Use:   "delete",
	Short: "删除证书",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("create server certificate")
		_ = modules.DeleteCertficate(global.CertIDs)
	},
}

func init() {
	CertCmdDelete.Flags().StringVarP(&global.CertIDs, "sc", "", "", `server certificates to delete, multiple ex "sc-123,sc-234"`)
}
