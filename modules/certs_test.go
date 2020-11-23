package modules

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_BindCerts(t *testing.T) {

	t.Run("BindCerts", func(t *testing.T) {
		ok := BindCertsToLBListener("sc-y1cftotl", "lbl-kt80mzjc")

		NewWithT(t).Expect(ok).Should(BeTrue())
	})
}
