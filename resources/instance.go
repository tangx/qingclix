package resources

type InstanceRequest struct {
	ImageId       string          `yaml:"image_id,omitempty" json:"image_id,omitempty" url:"image_id,omitempty"`
	CPU           string          `yaml:"cpu,omitempty" json:"cpu,omitempty" url:"cpu,omitempty"`
	Memory        string          `yaml:"memory,omitempty" json:"memory,omitempty" url:"memory,omitempty"`
	InstanceClass string          `yaml:"instance_class,omitempty" json:"instance_class,omitempty" url:"instance_class,omitempty"`
	OsDiskSize    string          `yaml:"os_disk_size,omitempty" json:"os_disk_size,omitempty" url:"os_disk_size,omitempty"`
	LoginMode     string          `yaml:"login_mode,omitempty" json:"login_mode,omitempty" url:"login_mode,omitempty"`
	LoginKeypair  string          `yaml:"login_keypair,omitempty" json:"login_keypair,omitempty" url:"login_keypair,omitempty"`
	InstanceName  string          `yaml:"instance_name,omitempty" json:"instance_name,omitempty" url:"instance_name,omitempty"`
	Zone          string          `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
	Vxnets        []string        `yaml:"vxnets,omitempty" json:"vxnets,omitempty" url:"vxnets.1,omitempty"`
	Contract      ContractRequest `yaml:"contract,omitempty" json:"contract,omitempty" url:"contract,omitempty"`
}

type RunInstancesResponse struct {
	Action    string   `json:"action"`
	Instances []string `json:"instances"`
	JobID     string   `json:"job_id"`
	RetCode   int64    `json:"ret_code"`
}
