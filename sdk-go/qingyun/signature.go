package qingyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

const ()

// Signature 签名机制
// https://docs.qingcloud.com/product/api/common/signature.html
func Signature(reqMethod string, platform string, param url.Values, secret string) string {

	// fmt.Println("param.Encode=", param.Encode())
	urls := strings.Replace(param.Encode(), "%25", "%", -1)
	source := reqMethod + "\n" + platform + "\n" + urls
	logrus.Debugf("source = %s\n", source)
	// fmt.Println(source)
	// return ShaHmac1(source, secret)
	return ShaHmac256(source, secret)
}

// ShaHmac1 计算 HMAC-SHA1 校验码
func ShaHmac1(source, secret string) string {
	key := []byte(secret)
	shaHmac := hmac.New(sha1.New, key)
	_, _ = shaHmac.Write([]byte(source))
	signedBytes := shaHmac.Sum(nil)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}

// ShaHmac1 计算 HMAC-SHA1 校验码
func ShaHmac256(source, secret string) string {
	key := []byte(secret)
	shaHmac := hmac.New(sha256.New, key)
	_, _ = shaHmac.Write([]byte(source))
	signedBytes := shaHmac.Sum(nil)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	signedString = strings.TrimSpace(signedString)
	signedString = strings.Replace(signedString, " ", "+", -1)
	signedString = url.QueryEscape(signedString)
	return signedString
}
