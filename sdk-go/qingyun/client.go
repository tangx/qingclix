package qingyun

import (
	"io/ioutil"

	"github.com/yunify/qingcloud-sdk-go/config"
	qservice "github.com/yunify/qingcloud-sdk-go/service"
	"gopkg.in/yaml.v2"
)

// Client config
type Client struct {
	QyAccessKeyID     string `yaml:"qy_access_key_id"`
	QySecretAccessKey string `yaml:"qy_secret_access_key"`
	Zone              string `yaml:"zone,omitempty"`
	// 青云官方 SDK cli
	qcli *qservice.QingCloudService
}

// New return
func New(secretID, secretKey, zone string) *Client {
	client := &Client{
		QyAccessKeyID:     secretID,
		QySecretAccessKey: secretKey,
		Zone:              zone,
	}

	client.initialOfficialClient()

	return client
}

// NewWithFile
func NewWithFile(file string) *Client {
	var user = Client{}

	fb, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(fb, &user)
	if err != nil {
		panic(err)
	}
	return New(user.QyAccessKeyID, user.QySecretAccessKey, user.Zone)
}

// 调用青云 SDK 返回客户端
func (cli *Client) initialOfficialClient() {

	if cli.qcli == nil {
		cfg, _ := config.New(cli.QyAccessKeyID, cli.QySecretAccessKey)
		qcli, _ := qservice.Init(cfg)

		cli.qcli = qcli
	}
}
