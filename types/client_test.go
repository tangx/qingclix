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
