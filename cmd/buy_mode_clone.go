package cmd

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/types"
)

func cloneMode() {

}

func clone(cli types.Client, instance string) (config ItemConfig, err error) {

	// insParams, volumes, contract := cloneInstance(cli, instance)

	// // clone instance
	// config.Instance = insParams

	// // clone volumes
	// for _, volume := range volumes {
	// 	volParams := cloneVolume(cli, volume)
	// 	config.Volumes = append()
	// }

	return

}
func cloneInstance(cli types.Client, instance string) (params types.RunInstancesRequest, volumes []string, contract string) {
	descParams := types.DescribeInstancesRequest{
		Instances: []string{instance},
	}
	descResp, err := cli.DescribeInstances(descParams)
	if err != nil {
		logrus.Errorf("%s", err)
		return
	}
	if descResp.TotalCount == 0 {
		msg := fmt.Sprintf("找不到指定主机: %s", instance)
		logrus.Infof("%s", errors.New(msg))
		return
	}

	// 遍历加入的网络
	target := descResp.InstanceSet[0]
	vxnets := []string{}
	for _, vxnet := range target.Vxnets {
		vxnets = append(vxnets, vxnet.VxnetID)
	}

	var keypair string

	if len(target.KeypairIDS) == 0 {
		keypair = ""
	} else {
		keypair = target.KeypairIDS[0]
	}

	// 初始化服务器请求变量
	params = types.RunInstancesRequest{
		ImageID:      target.Image.ImageID,
		InstanceType: target.InstanceType,
		InstanceName: target.InstanceName + "_clone",
		Zone:         target.ZoneID,
		Vxnets:       vxnets,
		LoginKeypair: keypair,
		LoginMode:    "keypair",
	}

	volumes = target.VolumeIDS
	contract = target.ReservedContract

	return
}
func cloneVolume(cli types.Client, volume string) (params types.CreateVolumesRequest) {
	return
}
