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
	RetCode         int    `json:"ret_code"`
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
	RetCode int           `json:"ret_code,omitempty"`
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

type DescribeReservedContractsRequest struct {
	// 合约状态: terminated(终止), cancelled(取消), pending(待支付/待审核), active(活跃), expired(过期), deleted(删除)
	Status []string `yaml:"status,omitempty" json:"status,omitempty" url:"status,omitempty,dotnumbered,numbered1"`
	// 用于筛选合约是否生效, 1表示获取已生效合约
	IsEffected int `yaml:"is_effected,omitempty" json:"is_effected,omitempty" url:"is_effected,omitempty"`
	// 合约处于待支付状态, pending(待审核), accepted(审核已通过待支付)
	TransitionStatus []string `yaml:"transition_status,omitempty" json:"transition_status,omitempty" url:"transition_status,omitempty,dotnumbered,numbered1"`

	Zone              string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
	Limit             int      `yaml:"limit,omitempty" json:"limit,omitempty" url:"limit,omitempty"`
	Offset            int      `yaml:"offset,omitempty" json:"offset,omitempty" url:"offset,omitempty"`
	ReservedContracts []string `yaml:"reserved_contracts,omitempty" json:"reserved_contracts,omitempty" url:"reserved_contracts,omitempty,dotnumbered,numbered1"`
	ResourceType      []string `yaml:"resource_type,omitempty" json:"resource_type,omitempty" url:"resource_type,omitempty,dotnumbered,numbered1"`
}

type DescribeReservedContractsResponse struct {
	Action              string                `json:"action"`
	TotalCount          int                   `json:"total_count"`
	ReservedContractSet []ReservedContractSet `json:"reserved_contract_set"`
	RetCode             int                   `json:"ret_code"`
}

type ReservedContractSet struct {
	LastApplyment     LastApplyment `json:"last_applyment"`
	AutoRenew         int           `json:"auto_renew"`
	Currency          string        `json:"currency"`
	CreateTime        string        `json:"create_time"`
	LeftMigrate       int           `json:"left_migrate"`
	ResourceLimit     int           `json:"resource_limit"`
	Fee               string        `json:"fee"`
	UserID            string        `json:"user_id"`
	ContractID        string        `json:"contract_id"`
	LeftUpgrade       int           `json:"left_upgrade"`
	Status            string        `json:"status"`
	AssociationMode   string        `json:"association_mode"`
	EffectTime        string        `json:"effect_time"`
	Description       string        `json:"description"`
	TransitionStatus  string        `json:"transition_status"`
	LastApplymentType string        `json:"last_applyment_type"`
	Entries           []interface{} `json:"entries"`
	ExpireTime        string        `json:"expire_time"`
	ZoneID            string        `json:"zone_id"`
	Months            int           `json:"months"`
	RootUserID        string        `json:"root_user_id"`
	ResourceType      string        `json:"resource_type"`
}

type LastApplyment struct {
	Status        string        `json:"status"`
	ApplymentType string        `json:"applyment_type"`
	Fee           string        `json:"fee"`
	UserID        string        `json:"user_id"`
	Description   string        `json:"description"`
	Months        int           `json:"months"`
	ApplymentID   string        `json:"applyment_id"`
	ConsoleID     string        `json:"console_id"`
	Currency      string        `json:"currency"`
	RootUserID    string        `json:"root_user_id"`
	CreateTime    string        `json:"create_time"`
	ResourceType  string        `json:"resource_type"`
	Entries       []interface{} `json:"entries"`
	Remarks       string        `json:"remarks"`
	StatusTime    string        `json:"status_time"`
	ContractID    string        `json:"contract_id"`
	ResourceLimit int           `json:"resource_limit"`
	ZoneID        string        `json:"zone_id"`
}

func (cli *Client) DescribeReservedContracts(params DescribeReservedContractsRequest) (resp DescribeReservedContractsResponse, err error) {
	err = cli.Get("DescribeReservedContracts", params, &resp)
	return
}

// 	解绑合约
type DissociateReservedContractRequest struct {
	Contract  string   `url:"contract,omitempty"`
	Resources []string `url:"resources,omitempty,dotnumbered,numbered1"`
}

// {"action":"DissociateReservedContractResponse","changed":["i-5xlu9gue"],"ret_code":0}

type DissociateReservedContractResponse struct {
	Action  string   `json:"action,omitempty"`
	Changed []string `json:"changed,omitempty"`
	RetCode int      `json:"ret_code,omitempty"`
	Message string   `json:"message,omitempty"`
}

func (cli *Client) DissociateReservedContract(params DissociateReservedContractRequest) (resp DissociateReservedContractResponse, err error) {
	err = cli.Get("DissociateReservedContract", params, &resp)
	return
}
