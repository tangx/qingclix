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

var (
	// 购买多台实例时， instance_name 增加后缀起点
	// example: InstancePostfixStartNumber=12
	// instance_name-12, instance_name-13
	PostfixStartNumber int
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&global.Verbose, "verbose", "v", 4, "logrus 日志等级。 0: Panic, 4: Info, 6: Trace. ")
	rootCmd.PersistentFlags().BoolVarP(&global.SkipContract, "skip_contract", "", false, "强制跳过合约购买过程。")
	rootCmd.PersistentFlags().IntVarP(&global.Count, "count", "c", 0, "设置购买数量, 0 为实例购买。 >1 为多实例增加后缀")
	rootCmd.PersistentFlags().IntVarP(&PostfixStartNumber, "postfix_start", "", 1, "设置多实例购买时的后缀的起始值")

	rootCmd.PersistentFlags().BoolVarP(&global.Dryrun, "dryrun", "", false, "不会真正购买。")

}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func QingclixSetLogLevel(level int) {
	logLevel := logrus.Level(level)
	logrus.SetLevel(logLevel)
}
