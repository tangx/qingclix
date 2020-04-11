package types

type CreateVolumesRequest struct {
	Size       int    `json:"size,omitempty" yaml:"size,omitempty" url:"size,omitempty"`
	VolumeName string `json:"volume_name,omitempty" yaml:"volume_name,omitempty" url:"volume_name,omitempty"`
	VolumeType int    `json:"volume_type,omitempty" yaml:"volume_type,omitempty" url:"volume_type,omitempty"`
	Count      int    `yaml:"count,omitempty" json:"count,omitempty" url:"count,omitempty"`
	Zone       string `json:"zone,omitempty" yaml:"zone,omitempty" url:"zone,omitempty"`
	Encryption string `yaml:"encryption,omitempty" json:"encryption,omitempty" url:"encryption,omitempty"`
	CipherALG  string `yaml:"cipher_alg,omitempty" json:"cipher_alg,omitempty" url:"cipher_alg,omitempty"`
}
type CreateVolumesResponse struct {
	Action  string   `url:"action,omitempty"`
	JobID   string   `url:"job_id,omitempty"`
	Volumes []string `url:"volumes,omitempty"`
	RetCode int      `url:"ret_code,omitempty"`
}

func (cli *Client) CreateVolumes(params CreateVolumesRequest) (resp CreateVolumesResponse, err error) {
	err = cli.Get("CreateVolumes", params, &resp)
	return
}

type AttachVolumesRequest struct {
	Volumes  []string `yaml:"volumes,omitempty" json:"volumes,omitempty" url:"volumes,omitempty,dotnumbered,numbered1"`
	Instance string   `yaml:"instance,omitempty" json:"instance,omitempty" url:"instance,omitempty"`
	Zone     string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}

type AttachVolumesResponse struct {
	Action  string `yaml:"action,omitempty" json:"action,omitempty"`
	JobID   string `yaml:"job_id,omitempty" json:"job_id,omitempty"`
	RetCode int    `yaml:"ret_code,omitempty" json:"ret_code,omitempty"`
}

func (cli *Client) AttachVolumes(params AttachVolumesRequest) (resp AttachVolumesResponse, err error) {
	err = cli.Get("AttachVolumes", params, &resp)
	return
}

type DetachVolumesRequest = AttachVolumesRequest
type DetachVolumesResponse = AttachVolumesResponse

func (cli *Client) DetachVolumes(params DetachVolumesRequest) (resp DetachVolumesResponse, err error) {
	err = cli.Get("DetachVolumes", params, &resp)
	return
}
