package qingyun

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_DescribeJobs(t *testing.T) {
	logrus.SetReportCaller(true)
	cli := NewWithFile(authFile)
	params := DescribeJobsRequest{
		Jobs:    []string{"j-gg5xgeec74l"},
		Verbose: 1,
	}
	resp, err := cli.DescribeJobs(params)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.TotalCount)

}
