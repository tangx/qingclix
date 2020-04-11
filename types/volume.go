package types

type VolumeRequest struct {
	Size       string `json:"size,omitempty" yaml:"size,omitempty" url:"size,omitempty"`
	VolumeName string `json:"volume_name,omitempty" yaml:"volume_name,omitempty" url:"volume_name,omitempty"`
	VolumeType string `json:"volume_type,omitempty" yaml:"volume_type,omitempty" url:"volume_type,omitempty"`
	Zone       string `json:"zone,omitempty" yaml:"zone,omitempty" url:"zone,omitempty"`
	Months     string `json:"months,omitempty" yaml:"months,omitempty" url:"months,omitempty"`
	AutoRenew  string `json:"auto_renew,omitempty" yaml:"auto_renew,omitempty" url:"auto_renew,omitempty"`
}

type AttachVolumesResquest struct {
	Volumes  []string `yaml:"volumes,omitempty" json:"volumes,omitempty" url:"volumes,omitempty"`
	Instance string   `yaml:"instance,omitempty" json:"instance,omitempty" url:"instance,omitempty"`
	Zone     string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}

type AttachVolumesResponse struct {
	Action  string `yaml:"action,omitempty" json:"action,omitempty"`
	JobID   string `yaml:"job_id,omitempty" json:"job_id,omitempty"`
	RetCode string `yaml:"ret_code,omitempty" json:"ret_code,omitempty"`
}
