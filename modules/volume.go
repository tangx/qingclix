package modules

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/sdk-go/qingyun"
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

func AttachVolume(instID string, volID string) (err error) {
	fmt.Printf("Attach Volume[%s] to Instance[%s] .", volID, instID)
	// defer fmt.Println("Err:", err)
	defer func() {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ok")
		}
	}()

	params := qingyun.AttachVolumesRequest{
		Instance: instID,
		Volumes:  []string{volID},
	}

	// check instance ok
	if err = CheckInstanceStatus(instID, "running"); err != nil {
		return err
	}

	// check volume ok
	if err = CheckVolumeAttachable(volID); err != nil {
		return err
	}

	// attach
	_, err = global.QingClix.AttachVolumes(params)
	if err != nil {
		logrus.Errorf("%s", err)
	}

	return nil
}

func CheckVolumeAttachable(volID string) (err error) {

	n := 60
	for i := 1; i <= n; i++ {
		if VolumeAttachable(volID) {
			break
		}

		if i == n {
			err = fmt.Errorf("%s Unattachable", volID)
			return err
		}
		fmt.Printf(".")
		time.Sleep(1 * time.Second)
	}

	return nil
}
