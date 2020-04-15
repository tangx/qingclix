package utils

import (
	"fmt"
	"testing"
)

func Test_HomeDir(t *testing.T) {
	fmt.Println(HomeDir())

	pathtmp := `/tmp/golang-gogogo/1/2`
	Mkdir(pathtmp)
}
