package e2e_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	"os/exec"
	"path/filepath"
	"time"
)

var _ = Describe("E2E", func() {

	Context("tanzubuilder", func() {

		When("Initializing a project", func() {

			var tmpDir string

			BeforeEach(func() {
				tmpDir = GinkgoT().TempDir()

				command := exec.Command(tanzubuilderCLI, "init")
				command.Dir = tmpDir

				session, err := Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session).WithTimeout(5 * time.Second).Should(Exit(0))
			})

			It("scaffolds a project", func() {
				Expect(filepath.Join(tmpDir, "Makefile")).To(BeARegularFile())
				Expect(filepath.Join(tmpDir, "Makefile_help.awk")).To(BeARegularFile())
			})

			When("Discovering all make targets", func() {

				It("Shows help for each make target", func() {
					command := exec.Command("make")
					command.Dir = tmpDir

					session, err := Start(command, GinkgoWriter, GinkgoWriter)
					Expect(err).NotTo(HaveOccurred())
					Eventually(session).WithTimeout(5 * time.Second).Should(Exit(0))

					Expect(session.Out).To(gbytes.Say("e2e"))
				})

			})

		})

	})

})
