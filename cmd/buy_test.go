package cmd

import (
	"testing"

	"github.com/tangx/qingclix/types"
)

func Test_ApplyLeaseAssociation(t *testing.T) {
	volumes := []string{"vol-rkn2wv8q"}
	params := types.ApplyReservedContractWithResourcesRequest{
		Zone:      "pek3d",
		AutoRenew: 1,
		Months:    1,
	}
	cli := types.Client{}
	applyLeaseAssociateContract(cli, volumes, params)

}
