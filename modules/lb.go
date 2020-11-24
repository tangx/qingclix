package modules

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/sdk-go/qingyun"
)

func UpdateLoadBalancers(lbIDs string) bool {
	lbs := strings.Split(lbIDs, ",")
	params := qingyun.UpdateLoadBalancersRequest{
		Loadbalancers: lbs,
	}
	resp, err := global.QingClix.UpdateLoadBalancers(params)
	if err != nil {
		logrus.Fatalf("retcode: %d, msg: %s", resp.RetCode, err.Error())
	}

	if resp.RetCode == 0 {
		logrus.Infof("update lb request is accepted, job_id is %s", resp.JobID)
		return true
	}

	return false
}
