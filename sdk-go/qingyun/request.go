package qingyun

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tangx/go-querystring/query"
)

const (
	// QINGCLOUD api server
	apiServer   = "api.qingcloud.com"
	apiPlatform = "/iaas/"
	apiVersion  = "1"

	reqProtocol      = "https"
	SignatureMethod  = "HmacSHA256"
	SignatureVersion = "1"
	timeFormat       = "2006-01-02T15:04:05Z"
)

// ErrorResponse is default error response
type ErrorResponse struct {
	RetCode int    `json:"ret_code,omitempty"`
	Message string `json:"message,omitempty"`
}

// 请求 alidns API
// 当返回 error 为 nil 的时候， errResp 一定为空结构体。 否则可以通过 errResp.Message 查看错误信息。
func (cli *Client) request(method, action string, param url.Values, body io.Reader, respInfo interface{}) ([]byte, error) {
	if param == nil {
		param = url.Values{}
	}

	param = pathEcscape(param)

	// 设置时区:
	//    https://blog.csdn.net/qq_26981997/article/details/53454606
	loc, _ := time.LoadLocation("") //参数就是解压文件的“目录”+“/”+“文件名”。
	//fmt.Println(time.Now().In(loc))
	//timeNow := time.Now().In(loc)
	//timeNow.Format("2006-01-02T15:04:05Z")
	timestamp := time.Now().In(loc).Format(timeFormat)

	// 阿里云服务器时间使用的是 UTC 时区。 中国时区+8
	// 会一直报错: Specified time stamp or date value is expired
	param.Set("time_stamp", timestamp)

	// common body
	// param.Set("Format", "JSON")
	param.Set("signature_method", SignatureMethod)
	param.Set("signature_version", SignatureVersion)
	param.Set("version", apiVersion)
	param.Set("access_key_id", cli.QyAccessKeyID)

	// ActionBody 请求是传入
	//param.Set("DomainName", "example.com")
	param.Set("action", action)

	// 判断参数是否带有 zone, 否则使用 config 里面默认默认值
	if zone := param.Get("zone"); zone == "" {
		param.Set("zone", cli.Zone)
	}

	// 获取签名
	// 注意: 阿里云对用户 key 签名有特殊说明
	//    https://help.aliyun.com/document_detail/29747.html?spm=a2c4g.11186623.6.619.57ad2846HCScB1
	signature := Signature(method, apiPlatform, param, cli.QySecretAccessKey)
	logrus.Debugf("signature = %s\n", signature)
	// 请求体中增加签名参数
	//param.Set("Signature", url.QueryEscape(signature))
	// param.Set("signature", signature)

	// 构建 url 请求地址
	// https://api.qingcloud.com/iaas/?access_key_id=QYACCESSKEYIDEXAMPLE&action=DescribeInstances&expires=2013-08-29T07%3A42%3A25Z&limit=20&signature_method=HmacSHA256&signature_version=1&status.1=running&time_stamp=2013-08-29T06%3A42%3A25Z&version=1&zone=pek3b&signature=ihPnXFgsg5yyqhDN2IejJ2%2Bbo89ABQ1UqFkyOdzRITY%3D
	urls := strings.Replace(param.Encode(), "%25", "%", -1)
	reqURL := reqProtocol + "://" + apiServer + apiPlatform + "?" + urls + "&signature=" + signature
	// fmt.Println(reqURL)
	logrus.Debugf("reqURL = %s\n", reqURL)

	req, err := http.NewRequest(method, reqURL, body)
	if err != nil {
		return nil, err
	}

	// 发起请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// 关闭请求
	defer resp.Body.Close()

	// 获取 body 内容
	return ioutil.ReadAll(resp.Body)

}

// requestGET 使用 GET 方法请求 API
func (cli *Client) requestGET(action string, param url.Values, respInfo interface{}) ([]byte, error) {
	return cli.request("GET", action, param, nil, respInfo)
}

// Do 开始执行请求
func (cli *Client) Do(action string, body map[string]string, optional map[string]string, respInfo interface{}) ([]byte, error) {
	param := url.Values{}
	for k, v := range body {
		param.Set(k, v)
	}
	for k, v := range optional {
		param.Set(k, v)
	}

	return cli.requestGET(action, param, respInfo)
}

// GetByMap 请求直接使用 map[string]string 传递所有调用请求。 具体请求参数查看青云对应 API 文档
func (cli *Client) GetByMap(action string, body map[string]string) ([]byte, error) {
	return cli.Do(action, body, nil, nil)
}

// GetByUrlValues 请求直接使用 url.Values 作为参数。 具体请求参数查看青云对应 API 文档
func (cli *Client) GetByUrlValues(action string, params url.Values) ([]byte, error) {
	return cli.requestGET(action, params, nil)
}

// func MethodGET
func (cli *Client) MethodGET(action string, params interface{}, ptrResp interface{}) error {

	urlValues, err := query.Values(params)
	if err != nil {
		return err
	}

	data, err := cli.requestGET(action, urlValues, ptrResp)
	// logrus.Debugf("%s", data)

	// http request error
	if err != nil {
		return fmt.Errorf("Http Request Error: %s", err)
	}

	// qingcloud api request error
	errMsg := ErrorResponse{}
	_ = json.Unmarshal(data, &errMsg)
	if errMsg.RetCode != 0 {
		s := fmt.Sprintf("%d: %s", errMsg.RetCode, errMsg.Message)
		return errors.New(s)
	}

	// unmarshal response data
	err = json.Unmarshal(data, ptrResp)
	if err != nil {
		return fmt.Errorf("Unmarshal response data Error: %s", err)
	}

	return nil
}

func pathEcscape(params url.Values) url.Values {
	/*
		https://docs.qingcloud.com/product/api/common/signature.html
		编码时空格要转换成 “%20” , 而不是 “+”
		resolv: https://www.jianshu.com/p/2ba7dda583b5
	*/
	p := url.Values{}
	for k, values := range params {
		var v2 []string
		for _, v := range values {
			v2 = append(v2, url.PathEscape(v))
		}
		p[k] = v2
	}

	return p
}
