package cmd

import (
	"testing"

	"github.com/tangx/qingclix/types"
)

func Test_ApplyLeaseAssociation(t *testing.T) {
	params := types.ApplyReservedContractWithResourcesRequest{
		Zone:      "pek3d",
		AutoRenew: 1,
		Months:    1,
		Resources: []string{"vol-rkn2wv8q"},
	}
	cli := types.Client{}
	ApplyLeaseAssociateContract(cli, params)

}
