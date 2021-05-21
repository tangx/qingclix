package qingyun

import "net/url"

// do 开始执行请求
func (cli *Client) do(action string, body map[string]string, optional map[string]string, respInfo interface{}) ([]byte, error) {
	param := url.Values{}
	for k, v := range body {
		param.Set(k, v)
	}
	for k, v := range optional {
		param.Set(k, v)
	}

	return cli.requestGET(action, param, respInfo)
}

// getByMap 请求直接使用 map[string]string 传递所有调用请求。 具体请求参数查看青云对应 API 文档
// 为了方便
//    在不通过 struct 传入参数请求的时候， 通过 requestGET 进行API 测试。
func (cli *Client) getByMap(action string, body map[string]string) ([]byte, error) {
	return cli.do(action, body, nil, nil)
}
