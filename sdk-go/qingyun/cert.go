package qingyun

type CreateServerCertificateRequest struct {
	ServerCertificateName string `json:"server_certificate_name,omitempty" url:"server_certificate_name,omitempty" yaml:"server_certificate_name,omitempty"`
	CertificateContent    string `json:"certificate_content,omitempty" url:"certificate_content,omitempty" yaml:"certificate_content,omitempty"`
	PrivateKey            string `json:"private_key,omitempty" url:"private_key,omitempty" yaml:"private_key,omitempty"`
	Zone                  string `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}

type CreateServerCertificateResponse struct {
	Action              string `json:"action,omitempty" url:"action,omitempty"`
	ServerCertificateID string `json:"server_certificate_id,omitempty" url:"server_certificate_id,omitempty"`
	RetCode             int    `json:"ret_code,omitempty" url:"ret_code,omitempty"`
}

func (cli *Client) CreateCertificate(params CreateServerCertificateRequest) (resp CreateServerCertificateResponse, err error) {
	err = cli.MethodGET("CreateServerCertificate", params, &resp)
	return
}

type DeleteServerCertificatesRequest struct {
	ServerCertificates []string `yaml:"server_certificates,omitempty" json:"server_certificates,omitempty" url:"server_certificates,omitempty,dotnumbered,numbered1"`
	Zone               string
}

type DeleteServerCertificatesResponse struct {
	Action  string `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	JobID   string `yaml:"job_id,omitempty" json:"job_id,omitempty" url:"job_id,omitempty"`
	RetCode int    `yaml:"ret_code,omitempty" json:"ret_code,omitempty" url:"ret_code,omitempty"`
}

func (cli *Client) DeleteCertificates(params DeleteServerCertificatesRequest) (resp DeleteServerCertificatesResponse, err error) {
	err = cli.MethodGET("DeleteServerCertificates", params, &resp)
	return
}

type AssociateCertsToLBListenerRequest struct {
	ServerCertificates   []string `yaml:"server_certificates,omitempty" json:"server_certificates,omitempty" url:"server_certificates,omitempty,dotnumbered,numbered1"`
	LoadbalancerListener string   `yaml:"loadbalancer_listener,omitempty" json:"loadbalancer_listener,omitempty" url:"loadbalancer_listener,omitempty"`
	Zone                 string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}

type AssociateCertsToLBListenerResponse struct {
	Action  string `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	Retcode int    `yaml:"retcode,omitempty" json:"retcode,omitempty" url:"retcode,omitempty"`
}

func (cli *Client) AssociateCertsToLBListener(params AssociateCertsToLBListenerRequest) (resp AssociateCertsToLBListenerResponse, err error) {
	err = cli.MethodGET("AssociateServerCertsToLBListener", params, &resp)
	return
}

type DissociateCertsFromLBListenerRequest struct {
	ServerCertificates   []string `yaml:"server_certificates,omitempty" json:"server_certificates,omitempty" url:"server_certificates,omitempty,dotnumbered,numbered1"`
	LoadbalancerListener string   `yaml:"loadbalancer_listener,omitempty" json:"loadbalancer_listener,omitempty" url:"loadbalancer_listener,omitempty"`
	Zone                 string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}

type DissociateCertsFromLBListenerResponse struct {
	Action  string `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	RetCode int    `yaml:"ret_code,omitempty" json:"ret_code,omitempty" url:"ret_code,omitempty"`
}

func (cli *Client) DissociateCertsFromLBListener(params DissociateCertsFromLBListenerRequest) (resp DissociateCertsFromLBListenerResponse, err error) {
	err = cli.MethodGET("DissociateServerCertsFromLBListener", params, &resp)
	return
}

type DescribeCertsRequest struct {
	ServerCertificates []string `yaml:"server_certificates,omitempty" json:"server_certificates,omitempty" url:"server_certificates,omitempty,dotnumbered,numbered1"`
	SearchWord         string   `yaml:"search_word,omitempty" json:"search_word,omitempty" url:"search_word,omitempty"`
	Verbose            int      `yaml:"verbose,omitempty" json:"verbose,omitempty" url:"verbose,omitempty"`
	Offset             int      `yaml:"offset,omitempty" json:"offset,omitempty" url:"offset,omitempty"`
	Limit              int      `yaml:"limit,omitempty" json:"limit,omitempty" url:"limit,omitempty"`
	Zone               string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
}

type DescribeCertsResponse struct {
	Action               string              `yaml:"action,omitempty" json:"action,omitempty" url:"action,omitempty"`
	TotalCount           int                 `yaml:"total_count,omitempty" json:"total_count,omitempty" url:"total_count,omitempty"`
	ServerCertificateSet []ServerCertificate `yaml:"server_certificate_set,omitempty" json:"server_certificate_set,omitempty" url:"server_certificate_set,omitempty"`
	RetCode              int                 `yaml:"ret_code,omitempty" json:"ret_code,omitempty" url:"ret_code,omitempty"`
}

func (cli *Client) DescribeCerts(params DescribeCertsRequest) (resp DescribeCertsResponse, err error) {
	err = cli.MethodGET("DescribeServerCertificates", params, &resp)
	return
}
