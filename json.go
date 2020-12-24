package encoding

import (
	"encoding/json"
)

// JsonEncoding 是对基于json的编解码器的抽象
type JsonEncoding interface {
	Encoding
}

func NewJsonEncoding() JsonEncoding {
	return newJsonEncoding()
}

// jsonEncoding 实现了 JsonEncoding 接口
type jsonEncoding struct{}

func newJsonEncoding() *jsonEncoding {
	return &jsonEncoding{}
}

func (j *jsonEncoding) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j *jsonEncoding) Decode(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

