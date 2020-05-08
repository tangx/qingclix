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
	configureCmd.AddCommand(CloneCmd)
}

var CloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		cloneMain()
	},
}

var (
	instance_target string
)

func init() {
	CloneCmd.Flags().StringVarP(&instance_target, "instance", "i", "", "target instance to clone")
}

func cloneMain() {
	if instance_target == "" {
		return
	}
	item := configure.CloneInstance(instance_target)

	config := configure.AddItem(item.Instance.InstanceName, item)
	configure.DumpConfig(config)

}
