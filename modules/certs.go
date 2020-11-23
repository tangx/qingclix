package modules

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingyun-sdk-go/qingyun"
)

func CreateCertficate(name string, keypath string, crtpath string) (id string) {
	key, err := ioutil.ReadFile(keypath)
	if err != nil {
		// panic(err)
		logrus.Fatalf("error: %s\n", err.Error())
	}

	crt, err := ioutil.ReadFile(crtpath)
	if err != nil {
		logrus.Fatalf("error: %s\n", err.Error())
	}

	params := qingyun.CreateServerCertificateRequest{
		ServerCertificateName: name,
		PrivateKey:            fmt.Sprintf("%s", key),
		CertificateContent:    fmt.Sprintf("%s", crt),
	}

	resp, err := global.QingClix.CreateCertificate(params)
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	fmt.Printf(`{"certid":"%s"}\n`, resp.ServerCertificateID)
	return resp.ServerCertificateID

}

// BindCertsToLBListener assicoate one or more certificate file to a LB listener
// certsIDs=sc-123456,sc-223456 split with comon (,)
func BindCertsToLBListener(certsIDs string, lblID string) bool {
	certs := strings.Split(certsIDs, ",")
	params := qingyun.AssociateCertsToLBListenerRequest{
		ServerCertificates:   certs,
		LoadbalancerListener: lblID,
	}

	resp, err := global.QingClix.AssociateCertsToLBListener(params)
	if err != nil {
		logrus.Fatalf("failed: %s", err.Error())
		return false
	}

	if resp.Retcode == 0 {
		logrus.Printf("success: binding certs(%s) to (%s)", certs, lblID)
		return true
	}
	return false

}
