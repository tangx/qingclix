package modules

import (
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