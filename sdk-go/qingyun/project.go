package qingyun

type AddProjectResourceItemsRequest struct {
	ProjectID string   `yaml:"project_id,omitempty" json:"project_id,omitempty" url:"project_id,omitempty"`
	Resources []string `yaml:"resources,omitempty" json:"resources,omitempty" url:"resources,omitempty,dotnumbered,numbered1"`
	Zone      string   `json:"zone,omitempty" yaml:"zone,omitempty" url:"zone,omitempty"`
}

type AddProjectResourceItemsResponse struct {
	Message     string   `json:"message,omitempty"`
	Action      string   `json:"action,omitempty"`
	ProjectID   string   `json:"project_id,omitempty"`
	ResourceIDs []string `json:"resource_ids,omitempty"`
	RetCode     int      `json:"ret_code,omitempty"`
	ZoneID      string   `json:"zone_id,omitempty"`
}

func (cli Client) AddProjectResourceItems(params AddProjectResourceItemsRequest) (resp AddProjectResourceItemsResponse, err error) {
	err = cli.MethodGET("AddProjectResourceItems", params, &resp)
	return
}

type DeleteProjectResourceItemsRequest struct {
	ProjectID []string `yaml:"project_id,omitempty" json:"project_id,omitempty" url:"project_id,omitempty,dotnumbered,numbered1"` // Required
	Resources []string `yaml:"resources,omitempty" json:"resources,omitempty" url:"resources,omitempty,dotnumbered,numbered1"`    // Required
	Zone      string   `json:"zone,omitempty" yaml:"zone,omitempty" url:"zone,omitempty"`
}

type DeleteProjectResourceItemsResponse struct {
	Message     string   `json:"message,omitempty"`
	Action      string   `json:"action,omitempty"`
	ProjectID   []string `json:"project_id,omitempty"`
	ResourceIDs []string `json:"resource_ids,omitempty"`
	RetCode     int      `json:"ret_code,omitempty"`
	ZoneID      string   `json:"zone_id,omitempty"`
}

func (cli *Client) DeleteProjectResourceItems(params DeleteProjectResourceItemsRequest) (resp DeleteProjectResourceItemsResponse, err error) {
	err = cli.MethodGET("DeleteProjectResourceItems", params, &resp)
	return
}
