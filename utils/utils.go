package utils

import (
	"log"
	"os"
	"os/user"

	"github.com/sirupsen/logrus"
)

func HomeDir() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return user.HomeDir
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Mkdir 创建目录，循环创建
func Mkdir(path string) (bool, error) {
	exist, err := PathExists(path)
	if err != nil {
		logrus.Errorln(err)
		return false, err
	}

	if exist {
		logrus.Debugf("%s 已存在", path)
		return true, nil
	}

	err = os.Mkdir(path, 0755)
	if err != nil {
		logrus.Errorln(err)
		return false, err
	}
	logrus.Debugf("创建路径 %s 成功", path)

	return true, nil
}

// MkdirAll 递归创建目录
func MkdirAll(path string) (bool, error) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		logrus.Errorln(err)
		return false, err
	}

	logrus.Debugf("创建路径 %s 成功", path)
	return true, nil
}

// Mkdirp : MkdirAll 的别名， 名字拼写上更靠近 `mkdir -p`
func Mkdirp(path string) (bool, error) {
	return MkdirAll(path)
}
