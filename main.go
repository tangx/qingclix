package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/cmd"
	"github.com/tangx/qingclix/global"
)

func main() {
	level := logrus.Level(global.Verbose)
	logrus.SetLevel(level)

	cmd.Execute()

}
