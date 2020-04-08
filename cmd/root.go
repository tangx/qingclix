package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qingclix",
	Short: "qingclix 命令行用来聚合青云控制台的复杂处理流程，简化日常操作",
	Long: `青云控制台的操作常用操作复杂，
  例如，新购机器、更换操作系统 等
  实现目标根据预设参数或配置，快速实现日常操作`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
