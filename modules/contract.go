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

func CreateContract(resource string, autorenew int, months int, zone string) (contractID string) {
	if autorenew == 1 {
		months = 1
	}

	params := qingyun.ApplyReservedContractWithResourcesRequest{
		Resources: []string{resource},
		Zone:      zone,
		Months:    months,
		AutoRenew: autorenew,
	}

	// apply
	contractID = ApplyContractWithResources(params)

	// lease
	ok := LeaseContract(contractID)
	logrus.Debugf("LeaseContract %t", ok)
	// if !ok then reapply()

	// associate
	ok = AssociateContract(contractID, params.Resources)
	logrus.Debugf("AssociateContract %t", ok)
	// if !ok then reapply()

	return contractID
}

func ApplyContractWithResources(params qingyun.ApplyReservedContractWithResourcesRequest) (contractID string) {

	resp, err := global.QingClix.ApplyReservedContractWithResources(params)
	if err != nil {
		logrus.Errorf("%s", err)
	}

	return resp.ContractID
}

func LeaseContract(id string) bool {
	params := qingyun.LeaseReservedContractRequest{
		Contract: id,
	}
	_, err := global.QingClix.LeaseReservedContract(params)
	if err != nil {
		logrus.Errorf("%s", err)
		return false
	}

	return true
}

func AssociateContract(contract string, resources []string) bool {
	params := qingyun.AssociateReservedContractRequest{
		Contract:  contract,
		Resources: resources,
	}

	_, err := global.QingClix.AssociateReservedContract(params)
	if err != nil {
		logrus.Errorf("%s", err)
		return false
	}

	return true
}
