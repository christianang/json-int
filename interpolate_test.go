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
			interpolatedJSON, err := jsonint.Interpolate(`{
					"first": "((first-var))",
					"nested": {
						"second": "((second-var))"
					},
					"array": [
						"((array-var-1))",
						"some-regular-value",
						"((array-var-2))",
						{
							"complex-value": {
								"key": "((complex-array-var))"
							}
						},
						["((nested-array-var-1))", "((nested-array-var-2))"]
					]
				}`,
				map[string]string{
					"first-var":          "my-first-value",
					"second-var":         "my-second-value",
					"array-var-1":        "my-array-value-1",
					"array-var-2":        "my-array-value-2",
					"complex-array-var":  "my-complex-array-value",
					"nested-array-var-1": "my-nested-array-value-1",
					"nested-array-var-2": "my-nested-array-value-2",
				},
			)
			Expect(err).NotTo(HaveOccurred())

			Expect(interpolatedJSON).To(MatchJSON(`{
				"first": "my-first-value",
				"nested": {
					"second": "my-second-value"
				},
				"array": [
					"my-array-value-1",
					"some-regular-value",
					"my-array-value-2",
					{
						"complex-value": {
							"key": "my-complex-array-value"
						}
					},
					["my-nested-array-value-1", "my-nested-array-value-2"]
				]
			}`))
		})
	})
})
