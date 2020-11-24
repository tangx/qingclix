package modules

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingyun-sdk-go/qingyun"
)

// DescribeLBL return which LoadBalance is belong to
// only accept one lbl as argument at once
func DescribeLBL(lbl string) (resp qingyun.DescribeLoadBalancerListenersResponse) {
	lbls := strings.Split(lbl, "")
	params := qingyun.DescribeLoadBalancerListenersRequest{
		LoadbalancerListeners: lbls,
		Verbose:               1,
	}
	resp, err := global.QingClix.DescribeLoadBalancerListeners(params)
	if err != nil {
		logrus.Fatalf("%s", err.Error())
		return
	}

	return
}

func GetLbByLbl(lbl string) (lb string) {
	resp := DescribeLBL(lbl)
	if len(resp.LoadBalanceListenerSet) == 1 {
		lb = resp.LoadBalanceListenerSet[0].LoadbalancerID
		// logrus.Printf("https://console.qingcloud.com/%s/loadbalancers/loadbalancers/%s/", zone, lb)
		return
	}

	return
}
