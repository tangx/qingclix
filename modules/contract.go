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

func CreateContract() {
	// apply
	// lease
	// associate
}

func ApplyContractWithResources(params qingyun.ApplyReservedContractWithResourcesRequest) (contractID string) {

	if params.AutoRenew == 1 {
		params.Months = 1
	}

	resp, err := global.QingClix.ApplyReservedContractWithResources(params)
	if err != nil {
		logrus.Errorf("%s", err)
	}

	return resp.ContractID
}

func LeaseContract()     {}
func AssociateContract() {}
