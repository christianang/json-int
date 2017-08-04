package jsonint

import (
	"encoding/json"
	"fmt"
)

func Interpolate(data string, variables map[string]string) (string, error) {
	var dataJSON map[string]interface{}
	err := json.Unmarshal([]byte(data), &dataJSON)
	if err != nil {
		panic(err)
	}

	dataJSON, err = interpolate(dataJSON, variables)
	if err != nil {
		panic(err)
	}

	interpolatedData, err := json.Marshal(dataJSON)
	if err != nil {
		panic(err)
	}

	return string(interpolatedData), nil
}

func interpolate(dataJSON map[string]interface{}, variables map[string]string) (map[string]interface{}, error) {
	for key, value := range dataJSON {
		switch value.(type) {
		case string:
			for varKey, varValue := range variables {
				if value.(string) == fmt.Sprintf("((%s))", varKey) {
					dataJSON[key] = varValue
				}
			}
		case map[string]interface{}:
			var err error
			dataJSON[key], err = interpolate(value.(map[string]interface{}), variables)
			if err != nil {
				panic(err)
			}
		default:
			panic("value is an unknown type")
		}
	}

	return dataJSON, nil
}
