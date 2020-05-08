package configure

import (
	"github.com/tangx/qingclix/modules"
)

func CloneInstance(target string) (item ClixItem) {

	// var item = configure.ClixItem{}
	insResp := modules.DescInstance(target)

	if len(insResp.InstanceSet) != 1 {
		return
	}

	ins := insResp.InstanceSet[0]

	item.Instance = ItemInstance{
		InstanceClass: ins.InstanceClass,
		InstanceName:  ins.InstanceName + "_clone",
		ImageID:       ins.Image.ImageID,
		CPU:           ins.VcpusCurrent,
		Memory:        ins.MemoryCurrent,
		LoginMode:     "keypair",
		LoginKeypair:  ins.KeypairIDS[0],
		Zone:          ins.ZoneID,
		OsDiskSize:    ins.Extra.OSDiskSize,
	}

	for _, vxnet := range ins.Vxnets {
		item.Instance.Vxnets = append(item.Instance.Vxnets, vxnet.VxnetID)
	}

	// contract
	contractResp := modules.DescContract(ins.ReservedContract)
	if contractResp.TotalCount == 1 {
		contract := contractResp.ReservedContractSet[0]
		if contract.AutoRenew == 1 {
			item.Contract.AutoRenew = contract.AutoRenew
			item.Contract.Months = 1
		}
	}

	// volume
	for _, volID := range ins.VolumeIDS {
		resp := modules.DescVolume(volID)
		if resp.TotalCount != 1 {
			continue
		}
		vol := resp.DescribeVolumeSet[0]

		volume := ItemVolume{
			Size:       vol.Size,
			VolumeType: vol.VolumeType,
		}
		item.Volumes = append(item.Volumes, volume)
	}

	return item
}
