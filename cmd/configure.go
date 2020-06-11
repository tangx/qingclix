package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/cmd/configure"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "配置管理",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.AddCommand(configure.CloneCmd)
	configureCmd.AddCommand(configure.CustomizeCmd)
}
