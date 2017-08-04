package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	jsonTemplate = ``
)

var _ = Describe("json-int", func() {
	Context("when a json template file and variables arguments are provided", func() {
		var (
			jsonTemplateFile string
		)

		BeforeEach(func() {
			tempDir, err := ioutil.TempDir("", "")
			Expect(err).NotTo(HaveOccurred())

			err = ioutil.WriteFile(filepath.Join(tempDir, "template.json"), []byte(`{ "key": "((some-variable-name))" }`), os.ModePerm)
			Expect(err).NotTo(HaveOccurred())
		})

		PIt("returns json with variables interpolated into template", func() {
			command := exec.Command(pathToJsonInt,
				jsonTemplateFile,
				"-v", "some-variable-name=some-variable-value",
			)

			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session, "10s").Should(gexec.Exit(0))

			Expect(string(session.Out.Contents())).To(MatchJSON(`{"key": "some-variable-value"}`))
		})
	})
})
