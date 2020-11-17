package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "生成补全",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	rootCmd.GenBashCompletion(os.Stdout)
	// },
}

var completionCmdBash = &cobra.Command{
	Use:   "bash",
	Short: "生成 bash 补全",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = rootCmd.GenBashCompletion(os.Stdout)
	},
}

var completionCmdZsh = &cobra.Command{
	Use:   "zsh",
	Short: "生成 zsh 补全",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = rootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.AddCommand(completionCmdZsh)
	completionCmd.AddCommand(completionCmdBash)
}
