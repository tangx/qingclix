package modules

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/sdk-go/qingyun"
)

// DescribeLBL return which LoadBalance is belong to
// only accept one lbl as argument at once
func DescribeLBL(lbl string) (resp qingyun.DescribeLoadBalancerListenersResponse) {
	lbls := strings.Split(lbl, ",")
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

func GetLbByLblID(lbls string) string {

	m := make(map[string]int)
	resp := DescribeLBL(lbls)

	if len(resp.LoadBalanceListenerSet) > 0 {
		for _, lbl := range resp.LoadBalanceListenerSet {
			m[lbl.LoadbalancerID] = 0
		}
	}

	ml := []string{}
	for k := range m {
		ml = append(ml, k)
	}

	return strings.Join(ml, ",")
}
