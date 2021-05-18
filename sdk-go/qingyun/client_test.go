package qingyun

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

const (
	authFile = `/Users/tangxin/.qingcloud/config.yaml`
)

var (
	cli = NewWithFile(authFile)
)

func Test_Get(t *testing.T) {
	cli := NewWithFile(authFile)

	action := `DescribeZones`
	bodyByte, err := cli.requestGET(action, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", bodyByte)

}

// 启动实例
func Test_RunInstance(t *testing.T) {
	cli := NewWithFile(authFile)

	runInstancesData := map[string]string{
		// 操作系统镜像ID
		"image_id": "xenial6x64",
		// cpu
		"cpu": "2",
		// 内存 (GB)
		"memory": "4096",
		// 实例类型 202: 企业e2
		"instance_class": "202",
		// 系统盘大小(GB)
		"os_disk_size": "40",
		// 登录方式
		"login_mode":    "keypair",
		"login_keypair": "kp-2kodyll8",
		// 实例名称
		"instance_name": "qingyun-sdk-go-running-d",
		// 所属网络
		"vxnets.1": "vxnet-kj0thr5",
		// 所在可用区
		"zone": "pek3d",
	}
	// urlValue := cli.RunInstances(runInstancesData)

	bodyByte, _ := cli.do("RunInstances", runInstancesData, nil, nil)
	fmt.Printf("%s", bodyByte)
}

// Test_DescribeReservedContracts 列出所有预留约合
func Test_DescribeReservedContracts(t *testing.T) {
	cli := NewWithFile(authFile)

	bodyByte, _ := cli.do("DescribeReservedContracts", nil, nil, nil)
	fmt.Printf("%s\n", bodyByte)
}

// Test_DescribeReservedResources 根据预留合约ID 查询绑定信息
func Test_DescribeReservedResources(t *testing.T) {
	cli := NewWithFile(authFile)

	data := map[string]string{
		"contracts.1": "rc-kAtfWuFu",
		"zone":        "pek3c",
	}

	bodyByte, _ := cli.getByMap("DescribeReservedResources", data)
	fmt.Printf("%s\n", bodyByte)
}

func Test_DescribeVolumes(t *testing.T) {

	// logrus.SetLevel(logrus.DebugLevel)
	cli := NewWithFile(authFile)

	params := DescribeVolumesRequest{
		Volumes: []string{"vol-s958u8rj"},
	}

	resp, err := cli.DescribeVolumes(params)
	if err != nil {
		panic(err)
	}

	if resp.TotalCount == 1 {
		fmt.Println(resp)
	}
}

func Test_DescribeInstances(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	inst := "i-cjg5iepp"

	params := DescribeInstancesRequest{
		Instances: []string{inst},
	}
	resp, err := cli.DescribeInstances(params)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

}

func Test_DescribeZones(t *testing.T) {
	params := DescribeZonesRequest{}
	resp, err := cli.DescribeZones(params)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func Test_AttachVolume(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	vol := "vol-uydrnlax"
	inst := "i-sd29dfl2"

	params := AttachVolumesRequest{
		Volumes:  []string{vol},
		Instance: inst,
		Zone:     "pek2",
	}

	resp, err := cli.AttachVolumes(params)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(resp)
}
