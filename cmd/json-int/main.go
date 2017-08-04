package main

import (
	"fmt"

	jsonint "github.com/christianang/json-int"
)

func main() {
	json, err := jsonint.Interpolate("", map[string]string{})
	if err != nil {
		panic(err)
	}

	fmt.Println(json)
}
