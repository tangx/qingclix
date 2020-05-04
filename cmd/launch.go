package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

// launchCmd represents the launch command
var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "购买机器",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		LaunchMain()
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)
}

func LaunchMain() {

	chooseItem()
}

func chooseItem() ClixItem {
	config := LoadConfig()

	var opts []string
	for name := range config.Configs {
		opts = append(opts, name)
	}

	var qs = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Select{
				Message: "choose a item name to launch",
				Options: opts,
			},
		},
	}

	answer := struct {
		Name string
	}{}

	err := survey.Ask(qs, &answer)
	if err != nil {
		logrus.Fatalf("%s", err)
	}

	return config.Configs[answer.Name]
}

func buyInstance() {
	// todo
}

func buyVolume() {
	// todo
}

func attachVolume() {
	// todo
}

func leaseContract() {
	// todo
}
