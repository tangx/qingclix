package qingyun

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_LBUpdate(t *testing.T) {
	params := UpdateLoadBalancersRequest{
		Loadbalancers: []string{"lb-ckxsncca"},
	}
	cli := NewWithFile(authFile)
	resp, err := cli.UpdateLoadBalancers(params)
	if err != nil {
		panic(err)
	}

	NewWithT(t).Expect(resp.RetCode).To(Equal(0), "update LoadBalance success")
}
