package main

import (
	"encoding/json"
	"fmt"
)

func parseBase(json map[string]interface{}, fields ...string) (map[string]interface{}, error) {
	for i := 0; i < len(fields)-1; i++ {
		t, ok := json[fields[i]].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("could not read field in JSON: %v", fields[i])
		}
		json = t
	}
	return json, nil
}

func ParseJson[T any](json map[string]interface{}, fields ...string) (T, error) {
	var defvalue T
	if len(fields) > 1 {
		t, err := parseBase(json, fields...)
		if err != nil {
			return defvalue, err
		}
		json = t
	}
	t, ok := json[fields[len(fields)-1]].(T)
	if ok {
		return t, nil
	}
	return defvalue, fmt.Errorf(fmt.Sprintf("could not read field in JSON: %v", fields[len(fields)-1]))
}

func UnmarshalJson(msg string) (map[string]interface{}, error) {
	var parsed map[string]interface{}
	err := json.Unmarshal([]byte(msg), &parsed)
	return parsed, err
}
