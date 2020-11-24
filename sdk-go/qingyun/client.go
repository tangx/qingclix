package qingyun

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Client config
type Client struct {
	QyAccessKeyID     string `yaml:"qy_access_key_id"`
	QySecretAccessKey string `yaml:"qy_secret_access_key"`
	Zone              string `yaml:"zone,omitempty"`
}

// New return
func New(secretID, secretKey, zone string) *Client {
	return &Client{
		QyAccessKeyID:     secretID,
		QySecretAccessKey: secretKey,
		Zone:              zone,
	}
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
