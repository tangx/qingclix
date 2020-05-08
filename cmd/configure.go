package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/cmd/configure"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.AddCommand(configure.CloneCmd)
}
