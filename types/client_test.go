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
