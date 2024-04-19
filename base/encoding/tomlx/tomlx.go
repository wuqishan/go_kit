package tomlx

import (
	"bytes"
	"encoding/json"
	"github.com/BurntSushi/toml"
)

func Encode(v interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(buffer).Encode(v); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Decode(v []byte) (interface{}, error) {
	var result interface{}
	if err := toml.Unmarshal(v, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func DecodeTo(v []byte, result interface{}) (err error) {
	err = toml.Unmarshal(v, result)
	return err
}

func ToJson(v []byte) ([]byte, error) {
	if r, err := Decode(v); err != nil {
		return nil, err
	} else {
		return json.Marshal(r)
	}
}
