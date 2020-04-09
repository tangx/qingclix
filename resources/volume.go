package resources

type VolumeRequest struct {
	Size       string          `json:"size,omitempty" yaml:"size,omitempty" url:"size,omitempty"`
	VolumeName string          `json:"volume_name,omitempty" yaml:"volume_name,omitempty" url:"volume_name,omitempty"`
	VolumeType string          `json:"volume_type,omitempty" yaml:"volume_type,omitempty" url:"volume_type,omitempty"`
	Zone       string          `json:"zone,omitempty" yaml:"zone,omitempty" url:"zone,omitempty"`
	Months     string          `json:"months,omitempty" yaml:"months,omitempty" url:"months,omitempty"`
	AutoRenew  string          `json:"auto_renew,omitempty" yaml:"auto_renew,omitempty" url:"auto_renew,omitempty"`
	Contract   ContractRequest `json:"contract,omitempty" yaml:"contract,omitempty" url:"contract,omitempty"`
}

type VolumeRequest2 struct {
	Size       string          `json:"size,omitempty" yaml:"size,omitempty" url:"size,omitempty"`
	VolumeName string          `json:"volume_name,omitempty" yaml:"volume_name,omitempty" url:"volume_name,omitempty"`
	VolumeType string          `json:"volume_type,omitempty" yaml:"volume_type,omitempty" url:"volume_type,omitempty"`
	Zone       string          `json:"zone,omitempty" yaml:"zone,omitempty" url:"zone,omitempty"`
	Months     string          `json:"months,omitempty" yaml:"months,omitempty" url:"months,omitempty"`
	AutoRenew  string          `json:"auto_renew,omitempty" yaml:"auto_renew,omitempty" url:"auto_renew,omitempty"`
	Contract   ContractRequest `json:"contract,omitempty" yaml:"contract,omitempty" url:"contract,omitempty"`
}
