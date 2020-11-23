package modules

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tangx/qingclix/global"
	"github.com/tangx/qingyun-sdk-go/qingyun"
)

func CreateCertficate(name string, keypath string, crtpath string) (id string) {
	key, err := ioutil.ReadFile(keypath)
	if err != nil {
		// panic(err)
		log.Fatalf("error: %s\n", err.Error())
	}

	crt, err := ioutil.ReadFile(crtpath)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	params := qingyun.CreateServerCertificateRequest{
		ServerCertificateName: name,
		PrivateKey:            fmt.Sprintf("%s", key),
		CertificateContent:    fmt.Sprintf("%s", crt),
	}

	resp, err := global.QingClix.CreateCertificate(params)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf(`{"certid":"%s"}\n`, resp.ServerCertificateID)
	return resp.ServerCertificateID

}
