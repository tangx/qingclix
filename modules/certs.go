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

func DeleteCertficate(certs string) (ok bool) {
	scs := strings.Split(certs, ",")
	params := qingyun.DeleteServerCertificatesRequest{
		ServerCertificates: scs,
	}

	resp, err := global.QingClix.DeleteCertificates(params)
	if err != nil {
		logrus.Fatalln(err.Error())
		return false
	}

	if resp.RetCode == 0 {
		logrus.Infof("success: deleting certs jobid= %s", resp.JobID)
		return true
	}
	return false
}

// BindCertsToLBListener assicoate one or more certificate file to a LB listener
// certs=sc-123456,sc-223456 split with comon (,)
func BindCertsToLBListener(certs string, lbl string) bool {
	scs := strings.Split(certs, ",")
	params := qingyun.AssociateCertsToLBListenerRequest{
		ServerCertificates:   scs,
		LoadbalancerListener: lbl,
	}

	resp, err := global.QingClix.AssociateCertsToLBListener(params)
	if err != nil {
		logrus.Fatalf("failed: %s", err.Error())
		return false
	}

	if resp.Retcode == 0 {
		logrus.Printf("success: binding certs(%s) to lbl(%s)", scs, lbl)
		return true
	}
	return false

}

func UnbindCertsFromLBListener(certs string, lbl string) (ok bool) {
	scs := strings.Split(certs, ",")
	params := qingyun.DissociateCertsFromLBListenerRequest{
		ServerCertificates:   scs,
		LoadbalancerListener: lbl,
	}

	resp, err := global.QingClix.DissociateCertsFromLBListener(params)
	if err != nil {
		logrus.Fatalf("failed: %s", err.Error())
		return false
	}

	if resp.RetCode == 0 {
		logrus.Printf("success: unbinding certs(%s) from lbl(%s)", scs, lbl)
		return true
	}
	return false
}

// DescribeOneCertByID return Certs info
func DescribeOneCertByID(sc string) (resp qingyun.DescribeCertsResponse) {
	scs := []string{sc}

	params := qingyun.DescribeCertsRequest{
		ServerCertificates: scs,
		Verbose:            1,
	}

	resp, err := global.QingClix.DescribeCerts(params)
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	if len(resp.ServerCertificateSet) == 1 {
		return
	}
	return
}

func SearchCertByName(name string) map[string]string {
	params := qingyun.DescribeCertsRequest{
		SearchWord: name,
	}
	resp, err := global.QingClix.DescribeCerts(params)
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	certs := map[string]string{}
	for _, cert := range resp.ServerCertificateSet {
		name := cert.ServerCertificateName
		id := cert.ServerCertificateID

		fmt.Printf("%s : %s\n", id, name)
		certs[name] = id
	}
	return certs
}

func GetCertBindTo(sc string) (m map[string][]string) {

	m = make(map[string][]string)
	resp := DescribeOneCertByID(sc)

	if len(resp.ServerCertificateSet) != 1 {
		return
	}

	fmt.Printf("Cert: %s\n", sc)
	for _, lbl := range resp.ServerCertificateSet[0].LoadbalancerListeners {
		lblID := lbl.LoadbalancerListenerID
		lbID := lbl.LoadbalancerID

		if len(m[lbID]) == 0 {
			m[lbID] = []string{lblID}
		} else {
			m[lbID] = append(m[lbID], lblID)
		}
	}

	for lb, lbls := range m {
		fmt.Printf("    %s : %s\n", lb, strings.Join(lbls, " "))
	}
	fmt.Println("")

	return
}
