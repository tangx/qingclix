package modules

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingyun-sdk-go/qingyun"
)

func DescInstance(id string) qingyun.DescribeInstancesResponse {
	params := qingyun.DescribeInstancesRequest{
		Instances: []string{id},
		Verbose:   1,
	}

	resp, err := global.QingClix.DescribeInstances(params)
	if err != nil {
		logrus.Errorf("%s", err)
	}

	return resp
}
