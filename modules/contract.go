package modules

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/sdk-go/qingyun"
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
	// 强行略过合约购买
	if global.SkipContract {
		logrus.Info("Skip contract apply")
		return ""
	}

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

func DissociateContract(contract string, resources ...string) bool {
	if len(resources) == 0 {
		return false
	}

	params := qingyun.DissociateReservedContractRequest{
		Contract:  contract,
		Resources: resources,
	}

	_, err := global.QingClix.DissociateReservedContract(params)

	if err != nil {
		logrus.Errorf("%s", err)
		return false
	}

	return true
}

func TerminateContract(contract string, confirm bool) error {
	if !confirm {
		return fmt.Errorf("请确认释放")
	}

	params := qingyun.TerminateReservedContractRequest{
		ContractID: contract,
		IsConfirm:  1,
	}

	resp, err := global.QingClix.TerminateReservedContract(params)
	if err != nil {
		return err
	}

	if resp.RetCode != 0 {
		return fmt.Errorf("错误: %s, code: %d", resp.Message, resp.RetCode)
	}

	return nil
}
