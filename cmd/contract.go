package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// contractCmd represents the contract command
var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "合约管理",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(contractCmd)
}

func targetsFromString(str string) (ret []string) {
	for _, res := range strings.Split(str, ",") {
		res = strings.TrimSpace(res)
		if len(res) == 0 {
			continue
		}
		ret = append(ret, res)
	}
	return
}

func targetsFromFile(file string) (targets []string) {

	f, err := os.Open(file)
	if err != nil {
		logrus.Errorf("open file %s failed: %v", err)
		return
	}
	defer f.Close()

	buf := bufio.NewScanner(f)
	for buf.Scan() {
		line := buf.Text()

		ret := targetsFromString(line)
		targets = append(targets, ret...)
	}

	return
}

func isDissocateValidArgs(targets string) bool {
	if len(targets) == 0 && len(fileToDissocite) == 0 {
		return false
	}

	return true
}
