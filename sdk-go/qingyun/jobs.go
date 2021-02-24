package qingyun

import "time"

type DescribeJobsRequest struct {
	Jobs    []string `yaml:"jobs,omitempty" json:"jobs,omitempty" url:"jobs,omitempty,dotnumbered,numbered1"`
	Status  []string `yaml:"status,omitempty" json:"status,omitempty" url:"status,omitempty,dotnumbered,numbered1"`
	Verbose int      `yaml:"verbose,omitempty" json:"verbose,omitempty" url:"verbose,omitempty"`
	Offset  int      `yaml:"offset,omitempty" json:"offset,omitempty" url:"offset,omitempty"`
	Limit   int      `yaml:"limit,omitempty" json:"limit,omitempty" url:"limit,omitempty"`
	Zone    string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}

type DescribeJobsResponse struct {
	Action     string `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	JobSet     []Job  `yaml:"job_set,omitempty" json:"job_set,omitempty" url:"job_set,omitempty,dotnumbered,numbered1"`
	TotalCount int    `yaml:"total_count,omitempty" json:"total_count,omitempty" url:"total_count,omitempty"`
	RetCode    int    `yaml:"ret_code,omitempty" json:"ret_code,omitempty" url:"ret_code,omitempty"`
}

type Job struct {
	JobID       string    `yaml:"job_id,omitempty" json:"job_id,omitempty" url:"job_id,omitempty"`
	JobAction   string    `yaml:"job_action,omitempty" json:"job_action,omitempty" url:"job_action,omitempty"`
	CreateTime  time.Time `yaml:"create_time,omitempty" json:"create_time,omitempty" url:"create_time,omitempty"`
	ResourceIDs string    `yaml:"resource_ids,omitempty" json:"resource_ids,omitempty" url:"resource_ids,omitempty"`
	Owner       string    `yaml:"owner,omitempty" json:"owner,omitempty" url:"owner,omitempty"`
	Status      string    `yaml:"status,omitempty" json:"status,omitempty" url:"status,omitempty"`
	StatusTime  time.Time `yaml:"status_time,omitempty" json:"status_time,omitempty" url:"status_time,omitempty"`
}

func (cli *Client) DescribeJobs(params DescribeJobsRequest) (resp DescribeJobsResponse, err error) {
	err = cli.MethodGET("DescribeJobs", params, &resp)
	return
}
