package global

import (
	"fmt"
	"testing"

	"github.com/tangx/qingyun-sdk-go/qingyun"
)

func Test_Clix(t *testing.T) {
	params := qingyun.DescribeZonesRequest{}
	resp, err := QingClix.DescribeZones(params)

	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
