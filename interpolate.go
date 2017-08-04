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
		return "", fmt.Errorf("interpolate failed to unmarshal json: %s", err)
	}

	dataJSON, err = interpolateMap(dataJSON, variables)
	if err != nil {
		return "", fmt.Errorf("interpolate failed: %s", err)
	}

	interpolatedData, err := json.Marshal(dataJSON)
	if err != nil {
		return "", fmt.Errorf("interpolate failed to marshal json: %s", err)
	}

	return string(interpolatedData), nil
}

func interpolateMap(data map[string]interface{}, variables map[string]string) (map[string]interface{}, error) {
	for key, value := range data {
		value, err := interpolateValue(value, variables)
		if err != nil {
			return nil, err
		}
		data[key] = value
	}

	return data, nil
}

func interpolateSlice(data []interface{}, variables map[string]string) ([]interface{}, error) {
	var interpolatedData []interface{}
	for _, value := range data {
		value, err := interpolateValue(value, variables)
		if err != nil {
			return nil, err
		}
		interpolatedData = append(interpolatedData, value)
	}

	return interpolatedData, nil
}

func interpolateValue(value interface{}, variables map[string]string) (interface{}, error) {
	switch value.(type) {
	case string:
		for varKey, varValue := range variables {
			if value.(string) == fmt.Sprintf("((%s))", varKey) {
				value = varValue
			}
		}
	case map[string]interface{}:
		var err error
		value, err = interpolateMap(value.(map[string]interface{}), variables)
		if err != nil {
			return nil, err
		}
	case []interface{}:
		var err error
		value, err = interpolateSlice(value.([]interface{}), variables)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("value is an unknown type of %s", reflect.TypeOf(value))
	}

	return value, nil
}
