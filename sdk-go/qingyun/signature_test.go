package qingyun

import (
	"fmt"
	"net/url"
	"testing"
)

const (
	AKID = `QYACCESSKEYIDEXAMPLE`
	AKEY = `SECRETACCESSKEY`
)

func Test_Signature(t *testing.T) {
	data := map[string]string{
		"access_key_id":     AKID,
		"action":            "RunInstances",
		"count":             "1",
		"image_id":          "centos64x86a",
		"instance_name":     "demo",
		"instance_type":     "small_b",
		"login_mode":        "passwd",
		"login_passwd":      "QingCloud20130712",
		"signature_method":  "HmacSHA256",
		"signature_version": "1",
		"time_stamp":        "2013-08-27T14:30:10Z",
		"version":           "1",
		"vxnets.1":          "vxnet-0",
		"zone":              "pek1",
	}

	params := url.Values{}
	for k, v := range data {
		params.Add(k, v)
	}
	// params := data
	ret := Signature("GET", "/iaas/", params, AKEY)
	fmt.Println("ret=", ret)
}
