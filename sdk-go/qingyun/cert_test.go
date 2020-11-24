package qingyun

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

var crt = `-----BEGIN CERTIFICATE-----
MIIFaTCCBFGgAwIBAgISBAlzFRah8d5S6UKUd24JL97aMA0GCSqGSIb3DQEBCwUA
MEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD
adUe+x++R6ToIqzhnQcGvZJHnx6fAE48HINIe7gDOWiB+MZL5RenHsUueEoRYoUm
XTIWRjwr1i1HWyXO7x0SMucaM1YFN1Thrxm17nWDCFjlDygKNSGxdbNeN7VywK9t
BM078qoT4LwBgRv8Y1AuyHx5TZJAExAviW/1qQWIpbA+0XoS8vGwcvLTeAiD/xbH
VDoeP7oyTnZOokH3vDeR2O0iDq3hdga4GPbI2N8XThjayj66kv81/yRLJ5XJNrQZ
jTVUyKDASwWcsP7n27LLMhLjzcQ7cDNsG0s+coW53ZMXmYD2x4oCZTFlVmX8
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
MIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/
X4Po1QYz+3dszkDqMp4fklxBwXRsW10KXzPMTZ+sOPAveyxindmjkW8lGy+QsRlG
PfZ+G6Z6h7mjem0Y+iWlkYcV4PIWL1iwBi8saCbGS5jN2p8M+X+Q7UNKEkROb3N6
KOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==
-----END CERTIFICATE-----`

var key = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAp6L3pY5vGdwwEkvAparJTnuGGT+MjJh68q5o6Z2kOVrSQDi8
aTn0sALGFxk7Md65Wda3mD5ytWO+rJL+BMIPiTpzd2WaS85QfKZhqV3vbXOu3WIl
en/Ho8UzzMTYQUepePAI3cV91xSwNhx94TV32DM5/4JUHAX2duQwWw==
-----END RSA PRIVATE KEY-----`

func Test_CreateCert(t *testing.T) {
	// logrus.SetLevel(logrus.DebugLevel)

	cli := NewWithFile(authFile)

	ccReq := CreateServerCertificateRequest{
		ServerCertificateName: "cert-test2",
		CertificateContent:    crt,
		PrivateKey:            key,
		Zone:                  "pek3",
	}

	resp, err := cli.CreateCertificate(ccReq)
	if err != nil {
		panic(err)
	}
	spew.Dump(resp)
}

func Test_DescribeCertByName(t *testing.T) {
	params := DescribeCertsRequest{
		SearchWord: "tangxin",
		Verbose:    1,
	}

	resp, err := cli.DescribeCerts(params)
	if err != nil {
		panic(err)
	}
	spew.Dump(resp)
}

func Test_DescribeCertByID(t *testing.T) {
	params := DescribeCertsRequest{
		ServerCertificates: []string{"sc-0j6zpvru"},
		Verbose:            1,
	}

	resp, err := cli.DescribeCerts(params)
	if err != nil {
		panic(err)
	}
	spew.Dump(resp)
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}
