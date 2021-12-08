package modules

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingclix/sdk-go/qingyun"
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

func CheckInstanceStatus(instID string, status string) (err error) {
	for i := 1; i <= 60; i++ {
		if status != InstanceStatus(instID) {
			fmt.Printf(".")
			time.Sleep(1 * time.Second)
			continue
		}
		return nil
	}
	err = fmt.Errorf("Check status Failed or not Match")
	return err
}

// StopInstance 关机
func StopInstance(ins []string, force int) (jobid string) {
	if len(ins) == 0 {
		return
	}

	req := qingyun.StopInstancesRequest{
		Instances: ins,
		Force:     force,
	}
	resp, err := global.QingClix.StopInstances(req)
	if err != nil {
		logrus.Errorf("关机失败, %s", ins)
		return
	}

	return resp.JobId
}

type InstanceAction string

const (
	InstanceActionStart   InstanceAction = "InstanceActionStart"
	InstanceActionStop    InstanceAction = "InstanceActionStop"
	InstanceActionRestart InstanceAction = "InstanceActionRestart"
)

func StartOrRestartInstances(ins []string, action InstanceAction) (jobid string) {
	if len(ins) == 0 {
		return
	}

	req := qingyun.BaseActionInstancesRequest{
		Instances: ins,
	}

	resp := qingyun.ActionInstancesResponse{}
	var err error

	switch action {
	case InstanceActionStart:
		resp, err = global.QingClix.StartInstances(req)
	case InstanceActionRestart:
		resp, err = global.QingClix.RestartInstances(req)
	}

	if err != nil {
		logrus.Errorf("%s instance %s: failed:", action, ins)
	}

	return resp.JobId
}
