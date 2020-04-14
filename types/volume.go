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

type DescribeVolumesRequest struct {
	Volumes    []string `yaml:"volumes,omitempty" json:"volumes,omitempty" url:"volumes,omitempty,dotnumbered,numbered1"`
	VolumeType int      `yaml:"volume_type,omitempty" json:"volume_type,omitempty" url:"volume_type,omitempty"`
	Status     []string `yaml:"status,omitempty" json:"status,omitempty" url:"status,omitempty,dotnumbered,numbered1"`
	SearchWord string   `yaml:"search_word,omitempty" json:"search_word,omitempty" url:"search_word,omitempty"`
	Tags       []string `yaml:"tags,omitempty" json:"tags,omitempty" url:"tags,omitempty,dotnumbered,numbered1"`
	Verbose    int      `yaml:"verbose,omitempty" json:"verbose,omitempty" url:"verbose,omitempty"`
	Offset     int      `yaml:"offset,omitempty" json:"offset,omitempty" url:"offset,omitempty"`
	Limit      int      `yaml:"limit,omitempty" json:"limit,omitempty" url:"limit,omitempty"`
	Zone       string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}
type DescribeVolumesResponse struct {
	Action            string              `json:"action,omitempty"`
	TotalCount        int                 `json:"total_count,omitempty"`
	DescribeVolumeSet []DescribeVolumeSet `json:"volume_set,omitempty"`
	RetCode           int                 `json:"ret_code,omitempty"`
}

type DescribeVolumeSet struct {
	Size             int    `json:"size,omitempty"`
	Encryption       int    `json:"encryption,omitempty"`
	Repl             string `json:"repl,omitempty"`
	ReservedContract string `json:"reserved_contract,omitempty"`
	VolumeID         string `json:"volume_id,omitempty"`
	ZoneID           string `json:"zone_id,omitempty"`
	VolumeType       int    `json:"volume_type,omitempty"`
	VolumeName       string `json:"volume_name,omitempty"`
}

func (cli *Client) DescribeVolumes(params DescribeVolumesRequest) (resp DescribeVolumesResponse, err error) {
	err = cli.Get("DescribeVolumes", params, &resp)
	return
}
