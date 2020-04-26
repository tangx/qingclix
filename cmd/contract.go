package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/qingclix/types"
)

// contractCmd represents the detach command
var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		DetachMain()
	},
}

// contractCmdDetach represents the detach command
var contractCmdDetach = &cobra.Command{
	Use:   "detach",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		DetachMain()
	},
}
var (
	detach_instace_with_volume bool = false
	detach_instance            string
)

func init() {
	rootCmd.AddCommand(contractCmd)
}

func init() {
	contractCmd.AddCommand(contractCmdDetach)
	contractCmdDetach.Flags().StringVarP(&detach_instance, "instance", "i", "", "解绑实例合约")
	contractCmdDetach.Flags().BoolVarP(&detach_instace_with_volume, "with-volume", "", false, "同时解绑磁盘合约")
}

func DetachMain() {
	cli := types.Client{}

	if detach_instance != "" {
		DetachInstance(cli, detach_instance)
	}

}

func DetachInstance(cli types.Client, instance string) {
	_, volumes, contract := cloneInstance(cli, instance)

	if contract != "" {
		DetachResourceContract(cli, instance, contract)
	}

	if !detach_instace_with_volume {
		return
	}
	for _, vol := range volumes {
		resp, _ := DescribleVolume(cli, vol)
		contract := resp.DescribeVolumeSet[0].ReservedContract
		if contract != "" {
			DetachResourceContract(cli, vol, contract)
		}
	}
}

func DescribleVolume(cli types.Client, volume string) (types.DescribeVolumesResponse, error) {
	params := types.DescribeVolumesRequest{
		Volumes: []string{volume},
	}
	return cli.DescribeVolumes(params)
}

func DetachResourceContract(cli types.Client, resource string, contract string) {

	params := types.DissociateReservedContractRequest{
		Contract:  contract,
		Resources: []string{resource},
	}
	_, err := cli.DissociateReservedContract(params)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Infof("资源[%s] 与 合约 [%s] 解绑成功", resource, contract)
}
