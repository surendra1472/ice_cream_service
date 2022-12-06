package utils

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalJSON_DataPresent(t *testing.T) {
	str := "success"
	var value *string
	var set bool
	strByte, _ := json.Marshal(str)
	err := UnmarshalJSON(strByte, new(string), &value, &set)
	assert.Nil(t, err)
	assert.True(t, set)
	assert.Equal(t, str, *value)
}

func TestUnmarshalJSON_NullData(t *testing.T) {
	var str *string
	var value *string
	var set bool
	strByte, _ := json.Marshal(str)
	err := UnmarshalJSON(strByte, new(string), &value, &set)
	assert.Nil(t, err)
	assert.Nil(t, value)
	assert.True(t, set)
}

func TestUnmarshalJSON_InvalidData(t *testing.T) {
	str := 1234
	var value *string
	var set bool
	strByte, _ := json.Marshal(str)
	err := UnmarshalJSON(strByte, new(string), &value, &set)
	assert.NotNil(t, err)
}
