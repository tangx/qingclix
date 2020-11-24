package cert

import (
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/modules"
)

var CertCmdSearch = &cobra.Command{
	Use:   "search",
	Short: "通过",
	Run: func(cmd *cobra.Command, args []string) {
		search()
	},
}

func init() {
	CertCmdSearch.Flags().StringVarP(&global.CertName, "name", "", "", "server certificate name or snippet")
}

func search() {
	_ = modules.SearchCertByName(global.CertName)
}
