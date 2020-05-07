package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "qingclix",
	Short: "一个青云简单的命令行工具",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&global.SkipContract, "skip-contract", "", false, "强制跳过合约购买")
}
