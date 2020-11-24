package qingyun

type LoadBalanceListener struct {
	ListenerProtocol         string        `json:"listener_protocol"`
	Forwardfor               int           `json:"forwardfor"`
	ListenerOption           int           `json:"listener_option"`
	HealthyCheckOption       string        `json:"healthy_check_option"`
	BalanceMode              string        `json:"balance_mode"`
	BackendProtocol          string        `json:"backend_protocol"`
	ListenerPortEnd          int           `json:"listener_port_end"`
	HealthyCheckMethod       string        `json:"healthy_check_method"`
	SessionSticky            string        `json:"session_sticky"`
	Scene                    int           `json:"scene"`
	LoadbalancerListenerName string        `json:"loadbalancer_listener_name"`
	Disabled                 int           `json:"disabled"`
	TunnelTimeout            int           `json:"tunnel_timeout"`
	CreateTime               string        `json:"create_time"`
	Timeout                  int           `json:"timeout"`
	Owner                    string        `json:"owner"`
	LoadbalancerListenerID   string        `json:"loadbalancer_listener_id"`
	LoadbalancerID           string        `json:"loadbalancer_id"`
	ListenerPort             int           `json:"listener_port"`
	WafDomainPolicies        []interface{} `json:"waf_domain_policies"`
	HTTPRedirectionCode      interface{}   `json:"http_redirection_code"`
	Backends                 []LblBackend  `json:"backends"`
}

type LblBackend struct {
	LoadbalancerBackendID   string `json:"loadbalancer_backend_id"`
	LoadbalancerBackendName string `json:"loadbalancer_backend_name"`
	Weight                  int    `json:"weight"`
	Port                    int    `json:"port"`
	ResourceID              string `json:"resource_id"`
	LoadbalancerListenerID  string `json:"loadbalancer_listener_id"`
	LoadbalancerID          string `json:"loadbalancer_id"`
	CreateTime              string `json:"create_time"`
}
