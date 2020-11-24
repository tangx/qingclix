package qingyun

type DescribeVxnetInstancesResponse struct {
	Action      string                                      `json:"action"`
	InstanceSet []DescribeVxnetInstancesResponseInstanceSet `json:"instance_set"`
	TotalCount  int                                         `json:"total_count"`
	RetCode     int                                         `json:"ret_code"`
}

type DescribeVxnetInstancesResponseInstanceSet struct {
	Ipv6Address      string                                         `json:"ipv6_address"`
	VxnetID          string                                         `json:"vxnet_id"`
	Aspoof           int                                            `json:"aspoof"`
	Sequence         int                                            `json:"sequence"`
	GraphicsPasswd   string                                         `json:"graphics_passwd"`
	Bandwidth        int                                            `json:"bandwidth"`
	CreateTime       string                                         `json:"create_time"`
	PrivateIP        string                                         `json:"private_ip"`
	Owner            string                                         `json:"owner"`
	NICID            string                                         `json:"nic_id"`
	MemoryCurrent    int                                            `json:"memory_current"`
	VcpusCurrent     int                                            `json:"vcpus_current"`
	NICName          string                                         `json:"nic_name"`
	SubCode          int                                            `json:"sub_code"`
	GraphicsProtocol string                                         `json:"graphics_protocol"`
	Platform         string                                         `json:"platform"`
	InstanceClass    int                                            `json:"instance_class"`
	Role             int                                            `json:"role"`
	StatusTime       string                                         `json:"status_time"`
	DHCPOptions      interface{}                                    `json:"dhcp_options"`
	Status           string                                         `json:"status"`
	Description      interface{}                                    `json:"description"`
	CPUTopology      string                                         `json:"cpu_topology"`
	Tags             []DescribeVxnetInstancesResponseInstanceSetTag `json:"tags"`
	InstanceID       string                                         `json:"instance_id"`
	TransitionStatus string                                         `json:"transition_status"`
	ImageID          string                                         `json:"image_id"`
	ZoneID           string                                         `json:"zone_id"`
	SriovNICType     int                                            `json:"sriov_nic_type"`
	InstanceName     string                                         `json:"instance_name"`
	InstanceType     string                                         `json:"instance_type"`
}

type DescribeVxnetInstancesResponseInstanceSetTag struct {
	Color    string `json:"color"`
	TagID    string `json:"tag_id"`
	TagValue string `json:"tag_value"`
	TagName  string `json:"tag_name"`
	Owner    string `json:"owner"`
	TagKey   string `json:"tag_key"`
}
