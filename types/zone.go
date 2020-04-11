package types

// type DescribeZonesRequest struct {
// 	Zone   []string `url:"zone,omitempty,dotnumbered,numbered1"`
// 	Status []string `url:"status,omitempty,dotnumbered,numbered1"`
// }

type DescribeZonesRequest struct {
	Zone   string `url:"zone,omitempty,dotnumbered,numbered1"`
	Status string `url:"status,omitempty,dotnumbered,numbered1"`
}

type DescribeZonesResponse struct {
	Action     string    `json:"action,omitempty"`
	TotalCount int       `json:"total_count,omitempty"`
	ZoneSet    []ZoneSet `json:"zone_set,omitempty"`
	RetCode    int       `json:"ret_code,omitempty"`
}

type ZoneSet struct {
	Status string `json:"status,omitempty"`
	ZoneID string `json:"zone_id,omitempty"`
}

func (cli *Client) DescribeZones(params DescribeZonesRequest) (resp DescribeZonesResponse, err error) {

	err = cli.Get("DescribeZones", params, &resp)

	return
}
