package cmd

import (
	"github.com/spf13/cobra"
)

// instanceCmd represents the instance command
var instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "服务器实例操作",
	Long: `针对服务器实例进行操作。
例如: 服务器配置管理、 服务器生命管理 等`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(instanceCmd)
}
