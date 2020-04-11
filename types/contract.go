package types

type ApplyReservedContractWithResourcesRequest struct {
	Zone      string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
	Resources []string `yaml:"resources,omitempty" json:"resources,omitempty" url:"resources,omitempty,dotnumbered,numbered1"`
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

func (cli *Client) ApplyReservedContractWithResources(params ApplyReservedContractWithResourcesRequest) (resp ApplyReservedContractWithResourcesResponse, err error) {
	err = cli.Get("ApplyReservedContractWithResources", params, &resp)
	return
}

type AssociateReservedContractRequest struct {
	Contract  string   `yaml:"contract,omitempty" json:"contract,omitempty" url:"contract,omitempty"`
	Resources []string `yaml:"resources,omitempty" json:"resources,omitempty" url:"resources,omitempty,dotnumbered,numbered1"`
}
type AssociateReservedContractResponse struct {
	Fail    []interface{} `json:"fail,omitempty"`
	Action  string        `json:"action,omitempty"`
	Success []string      `json:"success,omitempty"`
	RetCode int64         `json:"ret_code,omitempty"`
	Message string        `json:"message,omitempty"`
}

func (cli *Client) AssociateReservedContract(params AssociateReservedContractRequest) (resp AssociateReservedContractResponse, err error) {
	err = cli.Get("AssociateReservedContract", params, &resp)
	return
}

type LeaseReservedContractRequest struct {
	Contract         string `yaml:"contract,omitempty" json:"contract,omitempty" url:"contract,omitempty"`
	UnlimitedUpgrade int    `yaml:"unlimited_upgrade,omitempty" json:"unlimited_upgrade,omitempty" url:"unlimited_upgrade,omitempty"`
}
type LeaseReservedContractResponse struct {
	Action  string `json:"action,omitempty"`
	RetCode int    `json:"ret_code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (cli *Client) LeaseReservedContract(params LeaseReservedContractRequest) (resp LeaseReservedContractResponse, err error) {
	err = cli.Get("LeaseReservedContract", params, &resp)
	return
}
