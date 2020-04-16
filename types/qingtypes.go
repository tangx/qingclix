package types

import (
	"encoding/json"
)

type Qingtypes struct {
	InstanceType map[string]InstanceType `json:"instance_type,omitempty"`
	VolumeType   map[string]VolumeType   `json:"volume_type,omitempty"`
	ImageType    map[string]ImageType    `json:"image_type,omitempty"`
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
