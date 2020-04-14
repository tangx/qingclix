package cmd

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/types"
)

func cloneMode() {
	cli := types.Client{}

	if instance == "" {
		logrus.Errorln("没有指定克隆的目标机")
	}

	item, err := clone(cli, instance)
	if err != nil {
		panic(err)
	}

	// 直接购买
	if dump == "" {
		LaunchInstance(item)
	}

	// clone 配置
	if dump == "clone" {
		dump = item.Instance.InstanceName
	}
	item.Instance.InstanceName = dump
	dumpConfig(dump, item)

}

func dumpConfig(name string, item ItemConfig) {
	presetConfig := LoadPresetConfig()
	presetConfig.Configs[name] = item

	DumpPresetConfig(presetConfig)
}

func clone(cli types.Client, instance string) (config ItemConfig, err error) {

	insParams, volumes, contract := cloneInstance(cli, instance)

	// clone instance
	config.Instance = insParams

	// clone volumes
	volumesParams := cloneVolume(cli, volumes)
	config.Volumes = volumesParams

	// clone contract
	contractParams := cloneContract(cli, contract)
	config.Contract = contractParams

	return
}

func cloneInstance(cli types.Client, instance string) (params types.RunInstancesRequest, volumes []string, contract string) {
	descParams := types.DescribeInstancesRequest{
		Instances: []string{instance},
		Verbose:   1,
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
	logrus.Debugf("Keypair 数量: %d, keypair %s", len(target.KeypairIDS), target.KeypairIDS)

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

func cloneVolume(cli types.Client, volumes []string) (params []types.CreateVolumesRequest) {
	descParams := types.DescribeVolumesRequest{
		Volumes: volumes,
	}
	resp, err := cli.DescribeVolumes(descParams)
	if err != nil {
		logrus.Errorln(err)
	}

	for _, volume := range resp.DescribeVolumeSet {
		volParams := types.CreateVolumesRequest{
			Size:       volume.Size,
			VolumeType: volume.VolumeType,
		}
		params = append(params, volParams)
	}

	return
}

func cloneContract(cli types.Client, contract string) (params types.ApplyReservedContractWithResourcesRequest) {
	descParams := types.DescribeReservedContractsRequest{
		ReservedContracts: []string{contract},
	}
	resp, err := cli.DescribeReservedContracts(descParams)
	if err != nil {
		logrus.Errorln(err)
	}
	if resp.TotalCount == 0 {
		params.AutoRenew = 0
		params.Months = 0
		return
	}

	respContract := resp.ReservedContractSet[0]
	params.AutoRenew = respContract.AutoRenew
	params.Months = respContract.Months

	return
}
