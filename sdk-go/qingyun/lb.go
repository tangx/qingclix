package qingyun

type UpdateLoadBalancersRequest struct {
	Loadbalancers []string `yaml:"loadbalancers,omitempty" json:"loadbalancers,omitempty" url:"loadbalancers,omitempty,omitempty,dotnumbered,numbered1"`
	Zone          string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}

type UpdateLoadBalancersResponse struct {
	Action  string `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	JobID   string `yaml:"job_id,omitempty" json:"job_id,omitempty" url:"job_id,omitempty"`
	RetCode int    `yaml:"ret_code,omitempty" json:"ret_code,omitempty" url:"ret_code,omitempty"`
}

func (cli *Client) UpdateLoadBalancers(params UpdateLoadBalancersRequest) (resp UpdateLoadBalancersResponse, err error) {
	err = cli.MethodGET("UpdateLoadBalancers", params, &resp)
	return
}
