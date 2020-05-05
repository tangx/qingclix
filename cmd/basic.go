package cmd

type ClixConfig struct {
	Configs map[string]ClixItem `json:"configs"`
}

type ClixItem struct {
	Instance Instance `json:"instance"`
	Volumes  []Volume `json:"volumes"`
	Contract Contract `json:"contract"`
}

type Contract struct {
	Months    int `json:"months"`
	AutoRenew int `json:"auto_renew"`
}

type Instance struct {
	ImageID       string   `json:"image_id,omitempty" url:"image_id,omitempty"`
	CPU           int      `json:"cpu,omitempty" url:"cpu,omitempty"`
	Memory        int      `json:"memory,omitempty" url:"memory,omitempty"`
	InstanceClass int      `json:"instance_class,omitempty" url:"instance_class,omitempty"`
	LoginMode     string   `json:"login_mode,omitempty" url:"login_mode,omitempty"`
	LoginKeypair  string   `json:"login_keypair,omitempty" url:"login_keypair,omitempty"`
	InstanceName  string   `json:"instance_name,omitempty" url:"instance_name,omitempty"`
	Zone          string   `json:"zone,omitempty" url:"zone,omitempty"`
	Vxnets        []string `json:"vxnets,omitempty" url:"vxnets,omitempty"`
	OsDiskSize    int      `json:"os_disk_size,omitempty"`
}

type Volume struct {
	Size       int    `json:"size"`
	VolumeType int    `json:"volume_type"`
	VolumeName string `json:"volume_name"`
}
