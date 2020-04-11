package resources

import (
	"fmt"
	"testing"

	"github.com/tangx/qingclix/global"
)

const (

	// UBUNTU 16-32
	contract = "rc-77G6Y9Kt"
)

func Test_AssociateContract(t *testing.T) {
	cli := global.LoginQingyun()
	resources := []string{"i-x7ulv2i5"}
	contract := "rc-DRqvRvwg"
	resp := AssociateReservedContract(cli, contract, resources)
	fmt.Println(resp)
}

func Test_LeaseReservedContract(t *testing.T) {
	cli := global.LoginQingyun()
	// contract := "rc-3YtWtVGq"
	// unlimited_upgrade := 0

	resp := LeaseReservedContract(cli, contract)
	fmt.Println(resp)
}
