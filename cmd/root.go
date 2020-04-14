package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/global"
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

func init() {
	// rootCmd.Flags()(&mode, "mode", "m", "preset", ModeDesc)
	rootCmd.PersistentFlags().IntVarP(&global.Verbose, "verbose", "v", 4, "logrus 日志等级. 0: Panic, 4: Info, 6: Trace. ")
	rootCmd.PersistentFlags().BoolVarP(&global.SkipContract, "skip_contract", "", false, "true: 强制跳过合约购买过程; false 保持当前状态。")

}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func SetLogLevel(level int) {
	logLevel := logrus.Level(level)
	logrus.SetLevel(logLevel)
}
