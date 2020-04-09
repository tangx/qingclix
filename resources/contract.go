package resources

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
	Zone      string   `yaml:"zone,omitempty" json:"zone,omitempty"`
	Resources []string `yaml:"resources,omitempty" json:"resources,omitempty"`
	Months    int      `yaml:"months,omitempty" json:"months,omitempty"`
	AutoRenew int      `yaml:"auto_renew,omitempty" json:"auto_renew,omitempty"`
	User      string   `yaml:"user,omitempty" json:"user,omitempty"`
}
