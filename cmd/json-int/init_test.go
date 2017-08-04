package main

import (
	"testing"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	pathToJsonInt string
)

func TestJsonIntCLI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cmd/json-int")
}

var _ = BeforeSuite(func() {
	var err error
	pathToJsonInt, err = gexec.Build("github.com/christianang/json-int/cmd/json-int")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
