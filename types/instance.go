package types

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/tangx/go-querystring/query"
	"github.com/tangx/qingyun-sdk-go/qingyun"
)

type InstanceRequest struct {
	ImageId       string   `yaml:"image_id,omitempty" json:"image_id,omitempty" url:"image_id,omitempty"`
	CPU           string   `yaml:"cpu,omitempty" json:"cpu,omitempty" url:"cpu,omitempty"`
	Memory        string   `yaml:"memory,omitempty" json:"memory,omitempty" url:"memory,omitempty"`
	InstanceClass string   `yaml:"instance_class,omitempty" json:"instance_class,omitempty" url:"instance_class,omitempty"`
	OsDiskSize    string   `yaml:"os_disk_size,omitempty" json:"os_disk_size,omitempty" url:"os_disk_size,omitempty"`
	LoginMode     string   `yaml:"login_mode,omitempty" json:"login_mode,omitempty" url:"login_mode,omitempty"`
	LoginKeypair  string   `yaml:"login_keypair,omitempty" json:"login_keypair,omitempty" url:"login_keypair,omitempty"`
	InstanceName  string   `yaml:"instance_name,omitempty" json:"instance_name,omitempty" url:"instance_name,omitempty"`
	Zone          string   `yaml:"zone,omitempty" json:"zone,omitempty" url:"zone,omitempty"`
	Vxnets        []string `yaml:"vxnets,omitempty" json:"vxnets,omitempty" url:"vxnets,omitempty"`
}

type RunInstancesResponse struct {
	Action    string   `json:"action"`
	Instances []string `json:"instances"`
	JobID     string   `json:"job_id"`
	RetCode   int64    `json:"ret_code"`
	Message   string   `json:"message"`
}

func RunInstances(cli *qingyun.Client, config InstanceRequest) (resp RunInstancesResponse) {
	fmt.Printf("开始购买机器... ")

	action := "RunInstances"
	values, err := query.Values(config)
	if err != nil {
		logrus.Fatal("query.Values=", err)
	}

	body, err := cli.GetByUrlValues(action, values)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("%s\n", body)

	// var resp resources.RunInstancesResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		logrus.Fatal(err)
	}
	return
}

type DescribeInstancesRequest struct {
	Instances     []string
	ImageID       []string
	InstanceType  []string
	InstanceClass int
}
type DescribeInstancesResponse struct {
}
