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
	Instances     []string `yaml:"instances,omitempty" json:"instances,omitempty" url:"instances,omitempty,dotnumbered,numbered1"`
	ImageID       string   `yaml:"image_id,omitempty" json:"image_id,omitempty" url:"image_id,omitempty"`
	InstanceType  string   `yaml:"instance_type,omitempty" json:"instance_type,omitempty" url:"instance_type,omitempty,dotnumbered,numbered1"`
	InstanceClass int      `yaml:"instance_class,omitempty" json:"instance_class,omitempty" url:"instance_class,omitempty"`
	Status        []string `yaml:"status,omitempty" json:"status,omitempty" url:"status,omitempty,dotnumbered,numbered1"`
	Zone          string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}

type DescribeInstancesResponse struct {
	Action      string        `json:"action"`
	InstanceSet []InstanceSet `json:"instance_set"`
	RetCode     int64         `json:"ret_code"`
	TotalCount  int64         `json:"total_count"`
}

type InstanceSet struct {
	Vxnets           []Vxnet  `json:"vxnets"`
	MemoryCurrent    int64    `json:"memory_current"`
	VcpusCurrent     int64    `json:"vcpus_current"`
	Image            Image    `json:"image"`
	InstanceName     string   `json:"instance_name"`
	InstanceClass    int64    `json:"instance_class"`
	Status           string   `json:"status"`
	Description      string   `json:"description"`
	ReservedContract string   `json:"reserved_contract"`
	VolumeIDS        []string `json:"volume_ids"`
	ZoneID           string   `json:"zone_id"`
	InstanceID       string   `json:"instance_id"`
	InstanceType     string   `json:"instance_type"`
	Volumes          []Volume `json:"volumes"`
}

type Image struct {
	ImageID string `json:"image_id"`
}

type Volume struct {
	Device   string `json:"device"`
	VolumeID string `json:"volume_id"`
}

type Vxnet struct {
	Ipv6Address string `json:"ipv6_address"`
	VxnetType   int64  `json:"vxnet_type"`
	VxnetID     string `json:"vxnet_id"`
	VxnetName   string `json:"vxnet_name"`
	Role        int64  `json:"role"`
	PrivateIP   string `json:"private_ip"`
	NICID       string `json:"nic_id"`
}

type DescribeInstancesResponseInstanceSet struct{}

func (cli *Client) DescribeInstances(params DescribeInstancesRequest) (resp DescribeInstancesResponse, err error) {
	err = cli.Get("DescribeInstances", params, &resp)
	return
}
