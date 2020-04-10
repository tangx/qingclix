package resources

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"github.com/tangx/qingyun-sdk-go"
)

type ContractResponse struct {
	AutoRenew         int64         `json:"auto_renew,omitempty" yaml:"auto_renew,omitempty"`
	Currency          string        `json:"currency,omitempty" yaml:"currency,omitempty"`
	CreateTime        string        `json:"create_time,omitempty" yaml:"create_time,omitempty"`
	LeftMigrate       int64         `json:"left_migrate,omitempty" yaml:"left_migrate,omitempty"`
	ResourceLimit     int64         `json:"resource_limit,omitempty" yaml:"resource_limit,omitempty"`
	Fee               string        `json:"fee,omitempty" yaml:"fee,omitempty"`
	UserID            string        `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	ContractID        string        `json:"contract_id,omitempty" yaml:"contract_id,omitempty"`
	LeftUpgrade       int64         `json:"left_upgrade,omitempty" yaml:"left_upgrade,omitempty"`
	Status            string        `json:"status,omitempty" yaml:"status,omitempty"`
	AssociationMode   string        `json:"association_mode,omitempty" yaml:"association_mode,omitempty"`
	EffectTime        string        `json:"effect_time,omitempty" yaml:"effect_time,omitempty"`
	Description       string        `json:"description,omitempty" yaml:"description,omitempty"`
	TransitionStatus  string        `json:"transition_status,omitempty" yaml:"transition_status,omitempty"`
	LastApplymentType string        `json:"last_applyment_type,omitempty" yaml:"last_applyment_type,omitempty"`
	Entries           []interface{} `json:"entries,omitempty" yaml:"entries,omitempty"`
	ExpireTime        string        `json:"expire_time,omitempty" yaml:"expire_time,omitempty"`
	ZoneID            string        `json:"zone_id,omitempty" yaml:"zone_id,omitempty"`
	Months            int64         `json:"months,omitempty" yaml:"months,omitempty"`
	RootUserID        string        `json:"root_user_id,omitempty" yaml:"root_user_id,omitempty"`
	ResourceType      string        `json:"resource_type,omitempty" yaml:"resource_type,omitempty"`
}

type ContractRequest struct {
	Zone      string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
	Resources []string `yaml:"resources,omitempty" json:"resources,omitempty" url:"resources.1,omitempty"`
	Months    int      `yaml:"months,omitempty" json:"months,omitempty" url:"months,omitempty"`
	AutoRenew int      `yaml:"auto_renew,omitempty" json:"auto_renew,omitempty" url:"auto_renew,omitempty"`
	User      string   `yaml:"user,omitempty" json:"user,omitempty" url:"user,omitempty"`
}

type ApplyReservedContractWithResourcesResponse struct {
	Action          string `json:"action"`
	ApplymentStatus string `json:"applyment_status"`
	ContractID      string `json:"contract_id"`
	RetCode         int64  `json:"ret_code"`
}

type LeaseReservedContractRequest struct {
	Contract         string `json:"contract,omitempty" url:"contract,omitempty"`
	UnlimitedUpgrade string `json:"unlimited_upgrade,omitempty" url:"unlimited_upgrade,omitempty"`
}

// ApplyReservedContractWithResources 购买匹配资源的合约
func ApplyReservedContractWithResources(cli *qingyun.Client, contract ContractRequest) (resp ApplyReservedContractWithResourcesResponse) {
	fmt.Printf("开始购买合约... ")
	action := "ApplyReservedContractWithResources"

	values, err := query.Values(contract)
	if err != nil {
		logrus.Fatal()
	}
	body, err := cli.GetByUrlValues(action, values)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("%s\n", body)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		logrus.Fatal(err)
	}
	return resp
}

// type AssociateReservedContractResquest struct{}
type AssociateReservedContractResponse struct {
	Fail    []interface{} `json:"fail"`
	Action  string        `json:"action"`
	Success []string      `json:"success"`
	RetCode int64         `json:"ret_code"`
	Message string        `json:"message"`
}

type AssociateReservedContractRequest struct {
	Contract  string   `yaml:"contract,omitempty" json:"contract,omitempty" url:"contract,omitempty"`
	Resources []string `yaml:"resources,omitempty" json:"resources,omitempty" url:"resources.1,omitempty"`
}

// AssociateReservedContract 绑定合约和资源
func AssociateReservedContract(cli *qingyun.Client, contract string, resources []string) (resp AssociateReservedContractResponse) {
	fmt.Printf("绑定资源到合约... ")

	action := "AssociateReservedContract"
	params := AssociateReservedContractRequest{
		Contract:  contract,
		Resources: resources,
	}

	values, err := query.Values(params)
	if err != nil {
		logrus.Fatal()
	}
	body, err := cli.GetByUrlValues(action, values)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("%s\n", body)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		logrus.Fatal(err)
	}
	return resp
}
