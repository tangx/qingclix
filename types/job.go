package types

type DescribeJobsRequest struct {
	Jobs []string `yaml:"jobs,omitempty" json:"jobs,omitempty" url:"jobs,omitempty,dotnumbered,numbered1"`
	// Status 任务状态
	// pending, working, failed, successful
	Status  []string `yaml:"status,omitempty" json:"status,omitempty" url:"status,omitempty,dotnumbered,numbered1"`
	Verbose int      `yaml:"verbose,omitempty" json:"verbose,omitempty" url:"verbose,omitempty"`
	Offset  int      `yaml:"offset,omitempty" json:"offset,omitempty" url:"offset,omitempty"`
	Limit   int      `yaml:"limit,omitempty" json:"limit,omitempty" url:"limit,omitempty"`
	Zone    string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}

type DescribeJobsResponse struct {
	Action     string   `json:"action,omitempty"`
	TotalCount int      `json:"total_count,omitempty"`
	JobSet     []JobSet `json:"job_set,omitempty"`
	RetCode    int      `json:"ret_code,omitempty"`
}

type JobSet struct {
	Status      string          `json:"status,omitempty"`
	JobID       string          `json:"job_id,omitempty"`
	JobAction   string          `json:"job_action,omitempty"`
	CreateTime  string          `json:"create_time,omitempty"`
	Owner       string          `json:"owner,omitempty"`
	StatusTime  string          `json:"status_time,omitempty"`
	ErrorCodes  string          `json:"error_codes,omitempty"`
	ResourceIDS string          `json:"resource_ids,omitempty"`
	Resources   JobSetResources `json:"resources,omitempty"`
	Extras      interface{}     `json:"extras,omitempty"`
}

type JobSetResources struct {
	Instance string   `json:"instance,omitempty"`
	Volumes  []string `json:"volumes,omitempty"`
}

func (cli *Client) DescribeJobs(params DescribeJobsRequest) (resp DescribeJobsResponse, err error) {
	err = cli.Get("DescribeJobs", params, &resp)
	return
}
