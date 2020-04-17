package types

import (
	"encoding/json"
)

type Qingtypes struct {
	InstanceTypes []InstanceType `json:"instance_types,omitempty"`
	VolumeTypes   []VolumeType   `json:"volume_types,omitempty"`
	ImageTypes    []ImageType    `json:"image_types,omitempty"`
	Zones         []string       `json:"zones,omitempty"`
	Vxnets        []string       `json:"vxnets,omitempty"`
	Keypairs      []string       `json:"keypairs,omitempty"`
	Relationship  Relation       `json:"relation,omitempty"`
}

type ImageType struct {
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
	Desc  string `json:"desc,omitempty"`
}

type InstanceType struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Class int    `json:"class,omitempty"`
	Desc  string `json:"desc,omitempty"`
}

type Relation struct {
	Ship map[string][]int `json:"ship,omitempty"`
}

type VolumeType struct {
	Name string `json:"name,omitempty"`
	Type int    `json:"type,omitempty"`
	Desc string `json:"desc,omitempty"`
}

func LoadQingTypesString(data string) (qtypes Qingtypes, err error) {
	return LoadQingTypes([]byte(data))
}

func LoadQingTypes(data []byte) (qtypes Qingtypes, err error) {
	err = json.Unmarshal(data, &qtypes)
	if err != nil {
		return
	}
	return
}
