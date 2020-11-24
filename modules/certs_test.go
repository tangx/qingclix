package modules

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/gomega"
)

func Test_BindCerts(t *testing.T) {

	t.Run("BindCerts", func(t *testing.T) {
		ok := BindCertsToLBListener("sc-y1cftotl", "lbl-kt80mzjc")

		NewWithT(t).Expect(ok).Should(BeTrue())
	})
}
func Test_UnbindCerts(t *testing.T) {

	t.Run("BindCerts", func(t *testing.T) {
		ok := UnbindCertsFromLBListener("sc-y1cftotl", "lbl-kt80mzjc")

		NewWithT(t).Expect(ok).Should(BeTrue())
	})
}

func Test_DescribeOneCertByID(t *testing.T) {

	resp := DescribeOneCertByID("sc-0j6zpvru")
	spew.Dump(resp)
}

func Test_SearchCertByName(t *testing.T) {
	certs := SearchCertByName("wild")
	spew.Dump(certs)
}

func Test_GetCertBindToLbl(t *testing.T) {
	m := GetCertBindToLbl("sc-0j6zpvru")
	spew.Dump(m)
}
