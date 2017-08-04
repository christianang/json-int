package jsonint

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Interpolate(data string, variables map[string]string) (string, error) {
	var dataJSON map[string]interface{}
	err := json.Unmarshal([]byte(data), &dataJSON)
	if err != nil {
		panic(err)
	}

	dataJSON, err = interpolateMap(dataJSON, variables)
	if err != nil {
		panic(err)
	}

	interpolatedData, err := json.Marshal(dataJSON)
	if err != nil {
		panic(err)
	}

	return string(interpolatedData), nil
}

func interpolateMap(data map[string]interface{}, variables map[string]string) (map[string]interface{}, error) {
	for key, value := range data {
		switch value.(type) {
		case string:
			for varKey, varValue := range variables {
				if value.(string) == fmt.Sprintf("((%s))", varKey) {
					data[key] = varValue
				}
			}
		case map[string]interface{}:
			var err error
			data[key], err = interpolateMap(value.(map[string]interface{}), variables)
			if err != nil {
				panic(err)
			}
		case []interface{}:
			var err error
			data[key], err = interpolateSlice(value.([]interface{}), variables)
			if err != nil {
				panic(err)
			}
		default:
			panic(fmt.Sprintf("value is an unknown type %s", reflect.TypeOf(value)))
		}
	}

	return data, nil
}

func interpolateSlice(data []interface{}, variables map[string]string) ([]interface{}, error) {
	var interpolatedData []interface{}
	for _, value := range data {
		switch value.(type) {
		case string:
			variableExists := false
			for varKey, varValue := range variables {
				if value.(string) == fmt.Sprintf("((%s))", varKey) {
					interpolatedData = append(interpolatedData, varValue)
					variableExists = true
				}
			}
			if !variableExists {
				interpolatedData = append(interpolatedData, value)
			}
		default:
			panic(fmt.Sprintf("value is an unknown type %s", reflect.TypeOf(value)))
		}
	}

	return interpolatedData, nil
}
