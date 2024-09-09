package util

import (
	jsoniter "github.com/json-iterator/go"
)

var jsonNew = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(obj interface{}) (string, error) {
	bs, err := jsonNew.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func Unmarshal(jsonBytes []byte, obj interface{}) error {
	return jsonNew.Unmarshal(jsonBytes, obj)
}
