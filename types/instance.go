package types

type RunInstancesRequest struct {
	ImageID       string   `yaml:"image_id,omitempty" json:"image_id,omitempty" url:"image_id,omitempty"`
	CPU           int      `yaml:"cpu,omitempty" json:"cpu,omitempty" url:"cpu,omitempty"`
	Memory        int      `yaml:"memory,omitempty" json:"memory,omitempty" url:"memory,omitempty"`
	InstanceClass int      `yaml:"instance_class,omitempty" json:"instance_class,omitempty" url:"instance_class,omitempty"`
	OsDiskSize    string   `yaml:"os_disk_size,omitempty" json:"os_disk_size,omitempty" url:"os_disk_size,omitempty"`
	LoginMode     string   `yaml:"login_mode,omitempty" json:"login_mode,omitempty" url:"login_mode,omitempty"`
	LoginKeypair  string   `yaml:"login_keypair,omitempty" json:"login_keypair,omitempty" url:"login_keypair,omitempty"`
	InstanceName  string   `yaml:"instance_name,omitempty" json:"instance_name,omitempty" url:"instance_name,omitempty"`
	Zone          string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
	Vxnets        []string `yaml:"vxnets,omitempty" json:"vxnets,omitempty" url:"vxnets,omitempty,dotnumbered,numbered1"`
}

type RunInstancesResponse struct {
	Action    string   `json:"action"`
	Instances []string `json:"instances"`
	JobID     string   `json:"job_id"`
	RetCode   int      `json:"ret_code"`
	Message   string   `json:"message"`
}

func (cli *Client) RunInstances(params RunInstancesRequest) (resp RunInstancesResponse, err error) {
	err = cli.Get("RunInstances", params, &resp)
	return
}

type DescribeInstancesRequest struct {
	Instances     []string
	ImageID       []string
	InstanceType  []string
	InstanceClass int
}
type DescribeInstancesResponse struct {
}
