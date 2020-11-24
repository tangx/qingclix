package qingyun

import "time"

type ServerCertificate struct {
	ServerCertificateID         string                 `yaml:"server_certificate_id,omitempty" json:"server_certificate_id,omitempty" url:"server_certificate_id,omitempty"`
	ServerCertificateName       string                 `yaml:"server_certificate_name,omitempty" json:"server_certificate_name,omitempty" url:"server_certificate_name,omitempty"`
	PrivateKey                  string                 `yaml:"private_key,omitempty" json:"private_key,omitempty" url:"private_key,omitempty"`
	CertificateContent          string                 `yaml:"certificate_content,omitempty" json:"certificate_content,omitempty" url:"certificate_content,omitempty"`
	Description                 string                 `yaml:"description,omitempty" json:"description,omitempty" url:"description,omitempty"`
	CreateTime                  time.Time              `yaml:"create_time,omitempty" json:"create_time,omitempty" url:"create_time,omitempty"`
	EndTime                     time.Time              `yaml:"end_time,omitempty" json:"end_time,omitempty" url:"end_time,omitempty"`
	LoadbalancerListeners       []LoadBalanceListener  `yaml:"loadbalancer_listeners,omitempty" json:"loadbalancer_listeners,omitempty" url:"loadbalancer_listeners,omitempty,dotnumbered,numbered1"`
	ServerCertificateDependence []ServerCertDependence `yaml:"server_certificate_dependence,omitempty" json:"server_certificate_dependence,omitempty" url:"server_certificate_dependence,omitempty,dotnumbered,numbered1"`
}

type ServerCertDependence struct {
	ServerCertificateID    string `json:"server_certificate_id,omitempty" yaml:"server_certificate_id,omitempty" url:"server_certificate_id,omitempty"`
	Zone                   string `json:"zone,omitempty" yaml:"zone,omitempty" url:"zone,omitempty"`
	ExtraID                string `json:"extra_id,omitempty" yaml:"extra_id,omitempty" url:"extra_id,omitempty"`
	CreateTime             string `json:"create_time,omitempty" yaml:"create_time,omitempty" url:"create_time,omitempty"`
	ServiceType            string `json:"service_type,omitempty" yaml:"service_type,omitempty" url:"service_type,omitempty"`
	ServiceID              string `json:"service_id,omitempty" yaml:"service_id,omitempty" url:"service_id,omitempty"`
	ServerCertDependenceID string `json:"server_cert_dependence_id,omitempty" yaml:"server_cert_dependence_id,omitempty" url:"server_cert_dependence_id,omitempty"`
}
