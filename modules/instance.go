package modules

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingyun-sdk-go/qingyun"
)

func DescInstance(id string) qingyun.DescribeInstancesResponse {
	params := qingyun.DescribeInstancesRequest{
		Instances: []string{id},
		Verbose:   1,
	}

	resp, err := global.QingClix.DescribeInstances(params)
	if err != nil {
		logrus.Errorf("%s", err)
	}

	return resp
}

// RunInstance 够买一台服务器
// 没有 Count 参数传入，所以不能指定数量。
func RunInstance(params qingyun.RunInstancesRequest) (instID string) {
	fmt.Printf("Create Instance: ")
	resp, err := global.QingClix.RunInstances(params)
	if err != nil {
		logrus.Fatalf("%s", err)
	}

	if len(resp.Instances) == 0 {
		logrus.Fatalf("error: run instance failed")
	}

	instID = resp.Instances[0]
	fmt.Println(instID)

	return instID

}

func InstanceStatus(instanceID string) string {
	resp := DescInstance(instanceID)
	return resp.InstanceSet[0].Status

}

func CheckInstanceStatus(instanceID string, status string) error {
	for i := 1; i <= 20; i++ {
		if status != InstanceStatus(instanceID) {
			time.Sleep(1 * time.Second)
			continue
		}
		return nil
	}
	return fmt.Errorf("Check status Failed or not Match")
}
