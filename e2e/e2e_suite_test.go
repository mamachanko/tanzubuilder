package e2e_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var tanzubuilderCLI string

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}

var _ = BeforeSuite(func() {
	var err error
	tanzubuilderCLI, err = Build("github.com/mamachanko/tanzubuilder")
	DeferCleanup(CleanupBuildArtifacts)
	Expect(err).NotTo(HaveOccurred())
})
