package modules

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingyun-sdk-go/qingyun"
)

func DescVolume(id string) qingyun.DescribeVolumesResponse {
	params := qingyun.DescribeVolumesRequest{
		Volumes: []string{id},
	}

	resp, err := global.QingClix.DescribeVolumes(params)
	if err != nil {
		logrus.Errorf("%s", err)
	}
	return resp
}

func CreateVolume(params qingyun.CreateVolumesRequest) (volID string, err error) {

	fmt.Printf("Create Volume: ")
	resp, err := global.QingClix.CreateVolumes(params)
	if err != nil {
		return "", err
	}
	vol := resp.Volumes[0]
	fmt.Println(vol)

	return vol, nil
}

func VolumeStatus(volID string) (status string) {
	resp := DescVolume(volID)
	return resp.DescribeVolumeSet[0].Status

}

func VolumeAttachable(volID string) bool {
	resp := DescVolume(volID)
	return resp.DescribeVolumeSet[0].Attachable
}

func CheckVolumeStatus(volID string, status string) {
}
