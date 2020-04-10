package resources

import (
	"fmt"
	"testing"

	"github.com/tangx/qingclix/global"
)

func Test_AssociateContract(t *testing.T) {
	cli := global.LoginQingyun()
	resources := []string{"i-x7ulv2i5"}
	contract := "rc-DRqvRvwg"
	resp := AssociateReservedContract(cli, contract, resources)
	fmt.Println(resp)
}
