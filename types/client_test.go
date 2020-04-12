package types

import (
	"fmt"
	"testing"
)

func Test_DescribeZones(t *testing.T) {
	cli := Client{}
	params := DescribeZonesRequest{
		// Zone:   "pek3d",
		Status: "active",
	}
	resp, err := cli.DescribeZones(params)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.TotalCount, resp.RetCode, resp.Action)
}

func Test_RunInstances(t *testing.T) {
	cli := Client{}
	params := RunInstancesRequest{
		ImageID:       "xenial6x64",
		Vxnets:        []string{"vxnet-kj0thr5"},
		LoginMode:     "keypair",
		LoginKeypair:  "kp-2kodyll8",
		InstanceClass: 202,
		Memory:        2048,
		CPU:           1,
		Zone:          "pek3d",
	}

	resp, err := cli.RunInstances(params)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func Test_CreateVolume(t *testing.T) {
	cli := Client{}

	params := CreateVolumesRequest{
		Size:       100,
		VolumeName: "tangxin-test",
		VolumeType: 2,
		Zone:       "pek3d",
	}

	resp, err := cli.CreateVolumes(params)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func Test_AttachVolume(t *testing.T) {
	cli := Client{}
	params := AttachVolumesRequest{
		Volumes:  []string{"vol-uy2pywe2"},
		Instance: "i-x7ulv2i5",
		Zone:     "pek3d",
	}
	resp, err := cli.AttachVolumes(params)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func Test_DetachVolume(t *testing.T) {
	cli := Client{}
	params := DetachVolumesRequest{
		Volumes:  []string{"vol-uy2pywe2"},
		Instance: "i-x7ulv2i5",
		Zone:     "pek3d",
	}
	resp, err := cli.DetachVolumes(params)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func Test_DescribeInstances(t *testing.T) {
	cli := Client{}
	instance := "i-j33tcu6f"
	instance2 := "i-z625fhdq"
	status := "running"
	params := DescribeInstancesRequest{
		Instances: []string{instance, instance2},
		Status:    []string{status},
		Zone:      "pek3",
	}

	resp, err := cli.DescribeInstances(params)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
