// lbl for LoadBalanceListener
package qingyun

type DescribeLoadBalancerListenersRequest struct {
	LoadbalancerListeners []string `yaml:"loadbalancer_listeners,omitempty" json:"loadbalancer_listeners,omitempty" url:"loadbalancer_listeners,omitempty,dotnumbered,numbered1"`
	Loadbalancer          string   `yaml:"loadbalancer,omitempty" json:"loadbalancer,omitempty" url:"loadbalancer,omitempty"`
	Verbose               int      `yaml:"verbose,omitempty" json:"verbose,omitempty" url:"verbose,omitempty"`
	Offset                int      `yaml:"offset,omitempty" json:"offset,omitempty" url:"offset,omitempty"`
	Limit                 int      `yaml:"limit,omitempty" json:"limit,omitempty" url:"limit,omitempty"`
	Zone                  string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}

type DescribeLoadBalancerListenersResponse struct {
	Action                 string
	TotalCount             int
	LoadBalanceListenerSet []LoadBalanceListener
}

func (cli *Client) DescribeLoadBalancerListeners(params DescribeLoadBalancerListenersRequest) (resp DescribeLoadBalancerListenersResponse, err error) {
	err = cli.MethodGET("DescribeLoadBalancerListeners", params, &resp)
	return
}
