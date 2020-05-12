package integration_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Integration", func() {
	var (
		fizzbuzzCommand *exec.Cmd
		fizzbuzzBinary  string
		fizzbuzzArgs    []string
		session         *gexec.Session
	)

	JustBeforeEach(func() {
		var err error
		fizzbuzzBinary, err = gexec.Build("github.com/nikimanoledaki/fizzbuzz", "-mod=vendor")
		fizzbuzzCommand = exec.Command(fizzbuzzBinary, fizzbuzzArgs...)
		session, err = gexec.Start(fizzbuzzCommand, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		gexec.CleanupBuildArtifacts()
	})

	Context("when a comman line argument is received", func() {
		BeforeEach(func() {
			fizzbuzzArgs = []string{"3"}
		})

		It("prints the results", func() {
			Eventually(session.Out).Should(gbytes.Say("fizz\n"))
		})
	})

	Context("when several command line arguments are received", func() {
		BeforeEach(func() {
			fizzbuzzArgs = []string{"1", "2", "3", "4", "5"}
		})

		It("processes them all", func() {
			Eventually(session.Out).Should(gbytes.Say("1\n2\nfizz\n4\nbuzz\n"))
		})
	})

	Context("when the command line argument is not a number", func() {
		BeforeEach(func() {
			fizzbuzzArgs = []string{"!"}
		})

		It("it prints an erorr", func() {
			Eventually(session.Out).Should(gbytes.Say("arguments must be numbers\n"))
		})
	})

	Context("when there are no arguments at all", func() {
		BeforeEach(func() {
			fizzbuzzArgs = []string{}
		})

		It("it returns instructions", func() {
			Eventually(session.Out).Should(gbytes.Say("usage: fizzbuzz <number> \\[<number> <number> ...\\]\n"))
		})
	})
})
