package types

type InstanceVolume struct {
	Relationship []Relationship          `json:"relationship"`
	VolumeMap    map[string]ResourceDesc `json:"volume_map"`
	InstanceMap  map[string]ResourceDesc `json:"instance_map"`
}

type ResourceDesc struct {
	Class int    `json:"class"`
	Desc  string `json:"desc"`
}

type Relationship struct {
	InstanceTypes   []string `json:"instance_types"`
	InstanceClasses []int    `json:"instance_classes"`
	VolumeTypes     []int    `json:"volume_types"`
}
