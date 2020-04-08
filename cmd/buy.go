package cmd

import (
	"github.com/spf13/cobra"
)

// buyCmd represents the buy command
var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "根据预设信息购买机器",
	Long: `根据预设信息购买机器
  1. 根据 ~/.qingclix/preinstall.json 预设信息，直接选择购买对应的机器
  2. 使用 -i 进入交互界面，选择自定购买参数`,
	Run: func(cmd *cobra.Command, args []string) {
		launch()
	},
}

var interactive bool

func init() {
	rootCmd.AddCommand(buyCmd)
	buyCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "交互模式")
}

// launch 创建实例
func launch() {

	if interactive {
		// 交互模式
		interaciveMode()
	} else {

		preinstallMode()
	}

}

// interaciveMode 交互模式
func interaciveMode() {}

// preinstallMode 预设模式
func preinstallMode() {
	pre := loadPreinstallConfig()
	item := ChoosePreinstanllConfig(pre)
	launchPreinstall(item)
}
