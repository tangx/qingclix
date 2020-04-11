package types

import (
	"encoding/json"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/tangx/go-querystring/query"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingyun-sdk-go/qingyun"
)

// Client config
type Client struct {
	// QyAccessKeyID     string `yaml:"qy_access_key_id"`
	// QySecretAccessKey string `yaml:"qy_secret_access_key"`
	// Zone              string `yaml:"zone,omitempty"`
	Client *qingyun.Client
}

func (cli *Client) Login() {
	authfile := global.AuthFile
	if cli.Client == nil {
		cli.Client = qingyun.NewWithFile(authfile)
	}
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	RetCode int    `json:"ret_code,omitempty"`
}

// Get 请求青云 API 接口。 通过传入指针 &resp 的方式返回请求结果
func (cli *Client) Get(action string, params interface{}, resp interface{}) error {

	if cli.Client == nil {
		cli.Login()
	}

	values, err := query.Values(params)
	if err != nil {
		logrus.Fatal(err)
	}

	body, err := cli.Client.GetByUrlValues(action, values)
	if err != nil {
		logrus.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	// 判断请求是否错误
	errResp := ErrorResponse{}
	err = json.Unmarshal(body, &errResp)
	if err != nil {
		return err
	}
	if errResp.RetCode != 0 {
		return errors.New(string(body))
	}

	// 通过指针返回请求结果
	err = json.Unmarshal(body, resp)
	if err != nil {
		logrus.Fatal(err)
	}

	return nil
}
