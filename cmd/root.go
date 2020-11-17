package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "qingclix",
	Short: "一个青云简单的命令行工具",
	Long: `一个青云简单的命令行工具
version: ` + global.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initLogrusLevel()
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
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
	rootCmd.PersistentFlags().IntVarP(&global.Verbose, "verbose", "v", 4, "logrus 日志等级。 0: Panic, 4: Info, 6: Trace. ")
}

func initLogrusLevel() {
	level := logrus.Level(global.Verbose)
	logrus.SetLevel(level)
	logrus.Debugln(level)

}
