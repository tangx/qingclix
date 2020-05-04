package cmd

type ClixConfig struct {
	Configs map[string]ClixItem `json:"configs"`
}


type ClixItem struct {
	Instance Instance `json:"instance"`
	Volumes  []Volume `json:"volumes"` 
	Contract Contract `json:"contract"`
}

type Contract struct {
	Months    int `json:"months"`    
	AutoRenew int `json:"auto_renew"`
}

type Instance struct {
	ImageID       string   `json:"image_id"`      
	CPU           int    `json:"cpu"`           
	Memory        int    `json:"memory"`        
	InstanceClass int    `json:"instance_class"`
	LoginMode     string   `json:"login_mode"`    
	LoginKeypair  string   `json:"login_keypair"` 
	InstanceName  string   `json:"instance_name"` 
	Zone          string   `json:"zone"`          
	Vxnets        []string `json:"vxnets"`        
}

type Volume struct {
	Size       int `json:"size"`       
	VolumeType int `json:"volume_type"`
}