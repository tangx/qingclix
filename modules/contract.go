package modules

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingyun-sdk-go/qingyun"
)

func DescContract(id string) qingyun.DescribeReservedContractsResponse {
	params := qingyun.DescribeReservedContractsRequest{
		ReservedContracts: []string{id},
	}
	resp, err := global.QingClix.DescribeReservedContracts(params)
	if err != nil {
		logrus.Errorf("%s", err)
	}
	return resp
}
