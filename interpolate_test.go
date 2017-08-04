package jsonint_test

import (
	jsonint "github.com/christianang/json-int"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	jsonTemplate = ``
)

var _ = Describe("Interpolate", func() {
	Context("when a json template and variables are provided", func() {
		It("returns json with variables interpolated into template", func() {
			interpolatedJSON, err := jsonint.Interpolate(
				`{
					"first": "((first-var))",
					"nested": {
						"second": "((second-var))"
					}
				}`,
				map[string]string{
					"first-var":  "my-first-value",
					"second-var": "my-second-value",
				},
			)
			Expect(err).NotTo(HaveOccurred())

			Expect(interpolatedJSON).To(MatchJSON(`{
				"first": "my-first-value",
				"nested": {
					"second": "my-second-value"
				}
			}`))
		})
	})
})
