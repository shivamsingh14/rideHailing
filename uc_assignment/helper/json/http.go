package json

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

func Decode(body io.Reader, intf interface{}) (err error) {
	return jsoniter.ConfigFastest.NewDecoder(body).Decode(&intf)
}

func Marshal(file interface{}) ([]byte, error) {
	return jsoniter.ConfigFastest.Marshal(file)
}
