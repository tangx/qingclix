package types

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/tangx/go-querystring/query"
	"github.com/tangx/qingyun-sdk-go/qingyun"
)

type ContractRequest struct {
	Applyment          string              `yaml:"applyment,omitempty" json:"applyment,omitempty" url:"applyment,omitempty"`
	ApplymentType      string              `yaml:"applyment_type,omitempty" json:"applyment_type,omitempty" url:"applyment_type,omitempty"`
	AutoRenew          int                 `yaml:"auto_renew,omitempty" json:"auto_renew,omitempty" url:"auto_renew,omitempty"`
	Contract           string              `yaml:"contract,omitempty" json:"contract,omitempty" url:"contract,omitempty"`
	ContractId         string              `yaml:"contract_id,omitempty" json:"contract_id,omitempty" url:"contract_id,omitempty"`
	Contracts          []string            `yaml:"contracts,omitempty" json:"contracts,omitempty" url:"contracts,omitempty"`
	Currency           string              `yaml:"currency,omitempty" json:"currency,omitempty" url:"currency,omitempty"`
	Description        string              `yaml:"description,omitempty" json:"description,omitempty" url:"description,omitempty"`
	Entries            []map[string]string `yaml:"entries,omitempty" json:"entries,omitempty" url:"entries,omitempty"`
	IsEffected         int                 `yaml:"is_effected,omitempty" json:"is_effected,omitempty" url:"is_effected,omitempty"`
	LastApplymentType  string              `yaml:"last_applyment_type,omitempty" json:"last_applyment_type,omitempty" url:"last_applyment_type,omitempty"`
	Limit              int                 `yaml:"limit,omitempty" json:"limit,omitempty" url:"limit,omitempty"`
	Months             int                 `yaml:"months,omitempty" json:"months,omitempty" url:"months,omitempty"`
	Offset             int                 `yaml:"offset,omitempty" json:"offset,omitempty" url:"offset,omitempty"`
	ReservedApplyments []string            `yaml:"reserved_applyments,omitempty" json:"reserved_applyments,omitempty" url:"reserved_applyments,omitempty"`
	ReservedContracts  []string            `yaml:"reserved_contracts,omitempty" json:"reserved_contracts,omitempty" url:"reserved_contracts,omitempty"`
	ResourceType       string              `yaml:"resource_type,omitempty" json:"resource_type,omitempty" url:"resource_type,omitempty"`
	ResourceTypes      []string            `yaml:"resource_types,omitempty" json:"resource_types,omitempty" url:"resource_types,omitempty"`
	Resources          []string            `yaml:"resources,omitempty" json:"resources,omitempty" url:"resources,omitempty"`
	Reverse            int                 `yaml:"reverse,omitempty" json:"reverse,omitempty" url:"reverse,omitempty"`
	SearchWord         string              `yaml:"search_word,omitempty" json:"search_word,omitempty" url:"search_word,omitempty"`
	SortKey            string              `yaml:"sort_key,omitempty" json:"sort_key,omitempty" url:"sort_key,omitempty"`
	Status             string              `yaml:"status,omitempty" json:"status,omitempty" url:"status,omitempty"`
	ToZone             string              `yaml:"to_zone,omitempty" json:"to_zone,omitempty" url:"to_zone,omitempty"`
	TransitionStatus   string              `yaml:"transition_status,omitempty" json:"transition_status,omitempty" url:"transition_status,omitempty"`
	UnlimitedUpgrade   int                 `yaml:"unlimited_upgrade,omitempty" json:"unlimited_upgrade,omitempty" url:"unlimited_upgrade,omitempty"`
	User               string              `yaml:"user,omitempty" json:"user,omitempty" url:"user,omitempty"`
	Verbose            int                 `yaml:"verbose,omitempty" json:"verbose,omitempty" url:"verbose,omitempty"`
	Zone               string              `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}

type ApplyReservedContractWithResourcesResponse struct {
	Action          string `json:"action"`
	ApplymentStatus string `json:"applyment_status"`
	ContractID      string `json:"contract_id"`
	RetCode         int64  `json:"ret_code"`
}

// ApplyReservedContractWithResources 购买匹配资源的合约
func ApplyReservedContractWithResources(cli *qingyun.Client, contract ContractRequest) (resp ApplyReservedContractWithResourcesResponse) {
	fmt.Printf("开始购买合约... ")
	action := "ApplyReservedContractWithResources"

	fmt.Println(contract.Resources)
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

// AssociateReservedContract 绑定合约和资源
func AssociateReservedContract(cli *qingyun.Client, contract string, resources []string) (resp AssociateReservedContractResponse) {
	fmt.Printf("绑定资源到合约... ")

	action := "AssociateReservedContract"
	params := ContractRequest{
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

type LeaseReservedContractResponse struct {
	Action  string `json:"action,omitempty"`
	RetCode int    `json:"ret_code,omitempty"`
	Message string `json:"message,omitempty"`
}

// LeaseReservedContract 支付合约
func LeaseReservedContract(cli *qingyun.Client, contract string) (resp LeaseReservedContractResponse) {
	fmt.Printf("支付合约... ")

	action := "LeaseReservedContract"
	params := ContractRequest{
		Contract: contract,
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
