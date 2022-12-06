package utils

import (
	"encoding/json"
	"reflect"
)

func UnmarshalJSON(data []byte, valueObject interface{}, valuePtr interface{}, setPtr *bool) error {
	*setPtr = true
	if string(data) == "null" {
		return nil
	}
	err := json.Unmarshal(data, valueObject)
	if err != nil {
		return err
	}
	val := reflect.ValueOf(valuePtr)
	val.Elem().Set(reflect.ValueOf(valueObject))
	return nil
}
